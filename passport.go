package zkverifier_kit

import (
	"errors"
	"fmt"
	"maps"
	"os"
	"strconv"
	"time"

	val "github.com/go-ozzo/ozzo-validation/v4"
	zkptypes "github.com/iden3/go-rapidsnark/types"
	zkpverifier "github.com/iden3/go-rapidsnark/verifier"
	"github.com/rarimo/zkverifier-kit/root"
)

var ErrVerificationKeyRequired = errors.New("verification key is required")

// Verifier is a structure representing some instance for validation and verification zero knowledge proof
// generated by Rarimo system.
type Verifier struct {
	// verificationKey stores verification key content
	verificationKey []byte
	// opts has fields that must be validated before proof verification.
	opts VerifyOptions
}

// NewVerifier creates a new Verifier instance. VerificationKey is
// required to VerifyGroth16, usually you should just read it from file. Optional
// parameters will take part in proof verification on Verifier.VerifyProof call.
//
// If you provided WithVerificationKeyFile option, you can pass nil as the first arg.
func NewVerifier(verificationKey []byte, options ...VerifyOption) (*Verifier, error) {
	verifier := Verifier{
		verificationKey: verificationKey,
		opts:            mergeOptions(true, VerifyOptions{}, options...),
	}

	file := verifier.opts.verificationKeyFile
	if file == "" {
		if len(verificationKey) == 0 {
			return nil, ErrVerificationKeyRequired
		}
		return &verifier, nil
	}

	var err error
	verifier.verificationKey, err = os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read verification key from file %q: %w", file, err)
	}

	return &verifier, nil
}

// VerifyProof method verifies ZK proof and checks public signals. The public
// signals to validate are defined in the VerifyOption list. Firstly, you pass
// initial values to verify in NewVerifier. In case when custom values are
// required for different proofs, the options can be passed to VerifyProof, which
// override the initial ones.
//
// Filtered validation.Errors are always returned, unless this is internal error.
// You may use errors.As to assert whether it's validation or internal error.
func (v *Verifier) VerifyProof(proof zkptypes.ZKProof, options ...VerifyOption) error {
	v2 := Verifier{
		verificationKey: v.verificationKey,
		opts:            mergeOptions(false, v.opts, options...),
	}

	if err := v2.validatePubSignals(proof); err != nil {
		return err
	}

	if err := zkpverifier.VerifyGroth16(proof, v.verificationKey); err != nil {
		return val.Errors{
			"/proof": fmt.Errorf("groth16 verification failed: %w", err),
		}
	}

	return nil
}

func (v *Verifier) validatePubSignals(zkProof zkptypes.ZKProof) error {
	var (
		signals     = PubSignalGetter{ProofType: v.opts.proofType, Signals: zkProof.PubSignals}
		pubSigCount = PubSignalsCount(v.opts.proofType)
	)

	err := val.Errors{
		"zk_proof/proof":        val.Validate(zkProof.Proof, val.Required),
		"zk_proof/pub_signals":  val.Validate(zkProof.PubSignals, val.Required, val.Length(pubSigCount, pubSigCount)),
		"pub_signals/nullifier": val.Validate(signals.Get(Nullifier), val.Required),
	}.Filter()
	if err != nil {
		return err
	}

	if v.opts.proofType != PollParticipation {
		return v.validatePassportSignals(signals)
	}

	err = v.opts.voteVerifier.VerifyRoot(signals.Get(NullifiersTreeRoot))
	if !errors.Is(err, root.ErrInvalidRoot) {
		return err // internal error
	}

	return val.Errors{
		"participation_event_id": validateOnOptSet(signals.Get(ParticipationEventID), v.opts.partEventID, val.In(v.opts.partEventID)),
		"challenged_event_id":    validateOnOptSet(signals.Get(EventID), v.opts.eventID, val.In(v.opts.eventID)),
		"nullifiers_tree_root":   err,
	}.Filter()
}

func (v *Verifier) validatePassportSignals(signals PubSignalGetter) error {
	err := v.opts.passportVerifier.VerifyRoot(signals.Get(IdStateRoot))
	if err != nil && !errors.Is(err, root.ErrInvalidRoot) {
		return err
	}

	var (
		now       = time.Now().UTC()
		today     = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
		yesterday = today.AddDate(0, 0, -1)
		tomorrow  = today.AddDate(0, 0, 1)
	)

	all := val.Errors{
		"pub_signals/current_date": val.Validate(signals.Get(CurrentDate), val.When(
			v.opts.proofType == GeorgianPassport,
			val.Required,
			afterDate(yesterday),
			beforeDate(tomorrow),
		)),
		"pub_signals/personal_number_hash": val.Validate(signals.Get(PersonalNumberHash), val.When(
			v.opts.proofType == GeorgianPassport,
			val.Required,
		)),
		"pub_signals/id_state_root": err,
		"pub_signals/selector":      validateOnOptSet(signals.Get(Selector), v.opts.proofSelectorValue, val.In(v.opts.proofSelectorValue)),
		"pub_signals/event_id":      validateOnOptSet(signals.Get(EventID), v.opts.eventID, val.In(v.opts.eventID)),
		// upper bound is a date: the earlier it is, the higher the age
		"pub_signals/citizenship":   validateOnOptSet(decodeInt(signals.Get(Citizenship)), v.opts.citizenships, val.In(v.opts.citizenships...)),
		"pub_signals/event_data":    validateOnOptSet(signals.Get(EventData), v.opts.eventDataRule, v.opts.eventDataRule),
		"pub_signals/document_type": validateOnOptSet(decodeInt(signals.Get(DocumentType)), v.opts.documentType, val.In(v.opts.documentType)),
	}

	maps.Copy(all, v.validateBirthDate(signals))
	maps.Copy(all, v.validatePassportExpiration(signals))
	maps.Copy(all, v.validateIdentitiesInputs(signals))

	return all.Filter()
}

func (v *Verifier) validateBirthDate(signals PubSignalGetter) val.Errors {
	if v.opts.age == -1 {
		return nil
	}

	allowedBirthDate := time.Now().UTC().AddDate(-v.opts.age, 0, 0)
	return ORError(
		val.Validate(signals.Get(BirthDate), val.Required, beforeDate(allowedBirthDate)),
		val.Validate(signals.Get(BirthdateUpperBound), val.Required, equalDate(allowedBirthDate)),
		[2]string{"pub_signals/birth_date", "pub_signals/birth_date_upper_bound"},
	)
}

func (v *Verifier) validatePassportExpiration(signals PubSignalGetter) val.Errors {
	return val.Errors{
		"pub_signals/expiration_date_lower_bound": val.Validate(
			signals.Get(ExpirationDateLowerBound),
			val.When(!isEmptyZKDate(signals.Get(ExpirationDateLowerBound)), equalDate(time.Now().UTC())),
		),
		"pub_signals/expiration_date": val.Validate(
			signals.Get(ExpirationDate),
			val.When(!isEmptyZKDate(signals.Get(ExpirationDate)), afterDate(time.Now().UTC())),
		),
	}
}

// ZKP sets dates to 0 or 52983525027888 if date is not used or is not present in selector
func isEmptyZKDate(dateStr string) bool {
	return dateStr == "0" || dateStr == "52983525027888"
}

func (v *Verifier) validateIdentitiesInputs(signals PubSignalGetter) val.Errors {
	counter, err := strconv.ParseInt(signals.Get(IdentityCounterUpperBound), 10, 64)
	if err != nil {
		return val.Errors{"pub_signals/identity_counter_upper_bound": err}
	}

	// ZKP generates a timestamp upper bound as regular unix timestamp, so or time validation is not suitable here
	timestamp, err := strconv.ParseInt(signals.Get(TimestampUpperBound), 10, 64)
	if err != nil {
		return val.Errors{"pub_signals/timestamp_upper_bound": err}
	}

	return ORError(
		val.Validate(counter, val.When(
			v.opts.maxIdentitiesCount != -1,
			val.Max(v.opts.maxIdentitiesCount),
		)),
		validateOnOptSet(
			time.Unix(timestamp, 0),
			v.opts.maxIdentityCreationTimestamp,
			val.Max(v.opts.maxIdentityCreationTimestamp),
		),
		[2]string{"pub_signals/identity_counter_upper_bound", "pub_signals/timestamp_upper_bound"},
	)
}

func ORError(one, another error, fieldNames [2]string) val.Errors {
	// OR logic: at least one of the signals should be valid
	switch {
	case one != nil:
		return val.Errors{fieldNames[1]: another}
	case another != nil:
		return val.Errors{fieldNames[0]: one}
	default:
		return nil
	}
}

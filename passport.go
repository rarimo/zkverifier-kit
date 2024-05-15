package zkverifier_kit

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	val "github.com/go-ozzo/ozzo-validation/v4"
	zkptypes "github.com/iden3/go-rapidsnark/types"
	zkpverifier "github.com/iden3/go-rapidsnark/verifier"
	"github.com/rarimo/zkverifier-kit/identity"
)

type PubSignal int

// predefined values and positions for public inputs in zero knowledge proof. It
// may change depending on the proof and the values that it reveals.
const (
	Nullifier                 PubSignal = 0
	Citizenship               PubSignal = 6
	EventID                   PubSignal = 9
	EventData                 PubSignal = 10
	IdStateHash               PubSignal = 11
	Selector                  PubSignal = 12
	TimestampUpperBound       PubSignal = 14
	IdentityCounterUpperBound PubSignal = 16
	BirthdateUpperBound       PubSignal = 18
	ExpirationDateLowerBound  PubSignal = 19

	proofSelectorValue = "39"
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

// NewPassportVerifier creates a new Verifier instance. VerificationKey is
// required to VerifyGroth16, usually you should just read it from file. Optional
// parameters will take part in proof verification on Verifier.VerifyProof call.
//
// If you provided WithVerificationKeyFile option, you can pass nil as the first arg.
func NewPassportVerifier(verificationKey []byte, options ...VerifyOption) (*Verifier, error) {
	verifier := Verifier{
		verificationKey: verificationKey,
		opts:            mergeOptions(VerifyOptions{}, options...),
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

// VerifyProof method verifies iden3 ZK proof and checks public signals. The
// public signals to validate are defined in the VerifyOption list. Firstly, you
// pass initial values to verify in NewPassportVerifier. In case when custom
// values are required for different proofs, the options can be passed to
// VerifyProof, which override the initial ones.
func (v *Verifier) VerifyProof(proof zkptypes.ZKProof, options ...VerifyOption) error {
	v2 := Verifier{
		verificationKey: v.verificationKey,
		opts:            mergeOptions(v.opts, options...),
	}

	if err := v2.validate(proof); err != nil {
		return err
	}

	if err := zkpverifier.VerifyGroth16(proof, v.verificationKey); err != nil {
		return fmt.Errorf("groth16 verification failed: %w", err)
	}

	return nil
}

// validate is a helper method to validate public signals with values stored in opts field.
func (v *Verifier) validate(zkProof zkptypes.ZKProof) error {
	err := val.Errors{
		"zk_proof/proof":       val.Validate(zkProof.Proof, val.Required),
		"zk_proof/pub_signals": val.Validate(zkProof.PubSignals, val.Required, val.Length(21, 21)),
	}.Filter()
	if err != nil {
		return err
	}

	err = v.opts.rootVerifier.VerifyRoot(zkProof.PubSignals[IdStateHash])
	if errors.Is(err, identity.ErrContractCall) {
		return err
	}

	err = v.validateIdentitiesInputs(zkProof.PubSignals)
	if err != nil {
		return err
	}

	return val.Errors{
		// Required fields to validate
		"pub_signals/nullifier":                   val.Validate(zkProof.PubSignals[Nullifier], val.Required),
		"pub_signals/selector":                    val.Validate(zkProof.PubSignals[Selector], val.Required, val.In(proofSelectorValue)),
		"pub_signals/expiration_date_lower_bound": val.Validate(zkProof.PubSignals[ExpirationDateLowerBound], val.Required, afterDate(time.Now().UTC())),
		"pub_signals/id_state_hash":               err,

		// Configurable fields
		"pub_signals/event_id": val.Validate(zkProof.PubSignals[EventID], val.When(
			!val.IsEmpty(v.opts.eventID),
			val.Required,
			val.In(v.opts.eventID))),
		"pub_signals/birth_date_upper_bound": val.Validate(zkProof.PubSignals[BirthdateUpperBound], val.When(
			!val.IsEmpty(v.opts.age),
			val.Required,
			beforeDate(v.opts.age),
		)),
		"pub_signals/citizenship": val.Validate(decodeInt(zkProof.PubSignals[Citizenship]), val.When(
			!val.IsEmpty(v.opts.citizenships),
			val.Required,
			val.In(v.opts.citizenships...),
		)),
		"pub_signals/event_data": val.Validate(zkProof.PubSignals[EventData], val.When(
			!val.IsEmpty(v.opts.eventData),
			val.Required,
			val.In(v.opts.eventData),
		)),
	}.Filter()
}

func (v *Verifier) validateIdentitiesInputs(signals []string) error {
	counterUpperBound, err := strconv.ParseInt(signals[IdentityCounterUpperBound], 10, 64)
	if err != nil {
		return val.Errors{
			"pub_signals/identity_counter_upperbound": fmt.Errorf("parsing integer: %w", err),
		}.Filter()
	}

	counterErr := val.Errors{
		"pub_signals/identity_counter_upperbound": val.Validate(counterUpperBound, val.When(
			v.opts.maxIdentitiesCount != -1,
			val.Required,
			val.Max(v.opts.maxIdentitiesCount),
		)),
	}.Filter()

	timestampErr := val.Errors{
		"pub_signals/timestamp_upperbound": val.Validate(signals[TimestampUpperBound], val.When(
			!val.IsEmpty(v.opts.lastIdentityCreationTimestamp),
			val.Required,
			beforeDate(v.opts.lastIdentityCreationTimestamp),
		)),
	}.Filter()

	if counterErr != nil && timestampErr != nil {
		return errors.Join(counterErr, timestampErr)
	}

	if (counterErr != nil || timestampErr != nil) && !v.opts.allIdentitiesParamsSet {
		return errors.Join(counterErr, timestampErr)
	}

	return nil
}

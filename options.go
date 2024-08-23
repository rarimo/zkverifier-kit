package zkverifier_kit

import (
	"time"

	val "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/rarimo/zkverifier-kit/root"
)

// VerifyOptions structure that stores all fields that may be validated before proof verification.
// All elements must be able to have "zero" value in order to skip it during validation. For
// structure validation `github.com/go-ozzo/ozzo-validation/v4` is used, so IsEmpty method has to
// work correct with each field in order to have supposed logic.
type VerifyOptions struct {
	proofType proofType
	age       int
	// citizenships - array of interfaces (for more convenient usage during validation) that stores
	// all citizenships that accepted in proof. Under the hood, it is a string of Alpha-3 county codes,
	// described in the ISO 3166 international standard.
	citizenships  []interface{}
	eventDataRule val.Rule
	// eventID - unique identifier associated with a specific event or interaction within
	// the protocol execution, may be used to keep track of various steps or actions, this
	// id is a string with a big integer in decimals format
	eventID string
	// passportVerifier - verifies root in passport proof types
	passportVerifier root.Verifier
	// verificationKeyFile - stores verification key for proofs
	verificationKeyFile string
	// maxIdentitiesCount - maximum amount of reissued identities that user can have
	maxIdentitiesCount int64
	// maxIdentityCreationTimestamp - the upper bound of timestamp when user could create identities
	maxIdentityCreationTimestamp time.Time
	// proofSelectorValue - bit mask for selecting fields for verification
	proofSelectorValue string
	// documentType - provided document type name without UTF-8 encoding
	documentType string
	// partEventID - participation event ID in PollParticipation proof type
	partEventID string
	// voteVerifier - verifies root in PollParticipation proof type
	voteVerifier root.Verifier
	// skipExpirationCheck - specifies whether to check the expiration date.
	// Used for tests only
	skipExpirationCheck bool
}

// VerifyOption type alias for function that may add new values to VerifyOptions structure.
// It allows to create convenient methods With... that will add new value to the fields for
// that structure.
type VerifyOption func(*VerifyOptions)

// WithProofType select your proof type to use specific pub signals indexes.
// Default is GlobalPassport.
func WithProofType(t proofType) VerifyOption {
	return func(opts *VerifyOptions) {
		PubSignalsCount(t) // ensure proof type
		opts.proofType = t
	}
}

// WithAgeAbove adds new age check. It is an integer (e.g. 10, 18, 21) above which the person's
// age must be in proof.
func WithAgeAbove(age int) VerifyOption {
	return func(opts *VerifyOptions) {
		opts.age = age
	}
}

// WithCitizenships adds new available citizenship/s to prove that user is a resident of specified country.
// Function takes an arbitrary number of strings that consists from Alpha-3 county codes,
// described in the ISO 3166 international standard (e.g. "USA", "UKR", "TUR").
func WithCitizenships(citizenships ...string) VerifyOption {
	return func(opts *VerifyOptions) {
		opts.citizenships = make([]interface{}, len(citizenships))
		for i, ctz := range citizenships {
			opts.citizenships[i] = ctz
		}
	}
}

// WithEventData takes raw data for which the proof should be generated. This
// value format has to be validated before the signals validation because the kit
// checks ONLY the correspondence of these values. The raw value should be
// exactly in the same format as it was passed in the proof generation process
// and the length of the byte array is up to 31 bytes.
//
// Example: service takes Ethereum address, validates address format on their
// side, then converts address to bytes array and passes that bytes to the event
// data input in proof generation. After this precisely the same value has to be
// passed in the WithEventData function.
func WithEventData(raw []byte) VerifyOption {
	return func(opts *VerifyOptions) {
		opts.eventDataRule = eventData(raw)
	}
}

// WithEventID takes event identifier as a string that represents big number in a
// decimal format, i.e. it is compared with value from proof without conversions.
// This is used for PollParticipation flow too as verifier's (challenged) event ID.
func WithEventID(id string) VerifyOption {
	return func(opts *VerifyOptions) {
		opts.eventID = id
	}
}

// WithProofSelectorValue takes selector as a string that represents bit mask in a decimal format.
func WithProofSelectorValue(selector string) VerifyOption {
	return func(opts *VerifyOptions) {
		opts.proofSelectorValue = selector
	}
}

// WithPassportRootVerifier takes an abstract verifier that should verify IdStateRoot against the identity tree.
func WithPassportRootVerifier(v root.Verifier) VerifyOption {
	return func(opts *VerifyOptions) {
		opts.passportVerifier = v
	}
}

// WithVerificationKeyFile takes a string that represents the name of the file
// with verification key. The file is read on NewVerifier call. If you
// are providing this option along with the key argument, the latter will be
// overwritten by the read from file.
func WithVerificationKeyFile(name string) VerifyOption {
	return func(opts *VerifyOptions) {
		opts.verificationKeyFile = name
	}
}

// WithIdentitiesCounter takes maximum amount of identities that user can have
// during proof verification.
//
// On proof verification either this or WithIdentitiesCreationTimestampLimit pass.
func WithIdentitiesCounter(count int64) VerifyOption {
	return func(opts *VerifyOptions) {
		opts.maxIdentitiesCount = count
	}
}

// WithIdentitiesCreationTimestampLimit takes the upper bound for timestamp when
// user might create identities.
//
// On proof verification either this or WithIdentitiesCounter should pass.
func WithIdentitiesCreationTimestampLimit(unixTime int64) VerifyOption {
	return func(opts *VerifyOptions) {
		opts.maxIdentityCreationTimestamp = time.Unix(unixTime, 0)
	}
}

// WithDocumentType is used for DocumentType verification in GeorgianPassport proof
func WithDocumentType(docType string) VerifyOption {
	return func(opts *VerifyOptions) {
		opts.documentType = docType
	}
}

// WithPollParticipationEventID takes decimal eventID
func WithPollParticipationEventID(id string) VerifyOption {
	return func(opts *VerifyOptions) {
		opts.partEventID = id
	}
}

// WithPollRootVerifier takes an abstract verifier that should verify
// NullifiersTreeRoot against the state
func WithPollRootVerifier(v root.Verifier) VerifyOption {
	return func(opts *VerifyOptions) {
		opts.voteVerifier = v
	}
}

func withSkipExpirationCheck(check bool) VerifyOption {
	return func(opts *VerifyOptions) {
		opts.skipExpirationCheck = check
	}
}

// mergeOptions collects all parameters together and fills VerifyOptions struct
// with it, overwriting existing values
func mergeOptions(withDefaults bool, opts VerifyOptions, options ...VerifyOption) VerifyOptions {
	if withDefaults {
		opts.maxIdentitiesCount = -1
		opts.age = -1
		opts.passportVerifier = root.DisabledVerifier{}
		opts.proofType = GlobalPassport
	}

	for _, opt := range options {
		opt(&opts)
	}

	return opts
}

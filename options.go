package zkverifier_kit

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/rarimo/zkverifier-kit/identity"
)

// VerifyOptions structure that stores all fields that may be validated before proof verification.
// All elements must be able to have "zero" value in order to skip it during validation. For
// structure validation `github.com/go-ozzo/ozzo-validation/v4` is used, so IsEmpty method has to
// work correct with each field in order to have supposed logic.
type VerifyOptions struct {
	// externalID - is an external identifier with which the proof is associated. This value has
	// to be in a raw format (e.g. email, rarimo address), because this library will hash this value
	// with SHA256 hashing and compare with the one tha will be passed during proof verification.
	externalID string
	// age - a minimal age required to proof some statement.
	age time.Time
	// citizenships - array of interfaces (for more convenient usage during validation) that stores
	// all citizenships that accepted in proof. Under the hood, it is a string of Alpha-3 county codes,
	// described in the ISO 3166 international standard.
	citizenships []interface{}
	// eventData - is any data for which proof was generated, it has to be properly validated BEFORE the proof
	// verification because the SDK has no idea about the content inside, just checks if the values are the same.
	eventData interface{}
	// eventID - unique identifier associated with a specific event or interaction within
	// the protocol execution, may be used to keep track of various steps or actions, this
	// id is a string with a big integer in decimals format
	eventID string
	// rootVerifier - provider of identity root verification for IdStateHash
	rootVerifier IdentityRootVerifier
	// verificationKeyFile - stores verification key for proofs
	verificationKeyFile string
	// maxIdentitiesCount - maximum amount of identities that user can have. Default value is -1
	maxIdentitiesCount int64
	// lastIdentityCreationTimestamp - the upper timestamp when user might create their identities
	lastIdentityCreationTimestamp time.Time

	// helper option to check whether the all values for identities were set, this value is used for validation.
	allIdentitiesParamsSet bool
}

type IdentityRootVerifier interface {
	VerifyRoot(root string) error
}

// VerifyOption type alias for function that may add new values to VerifyOptions structure.
// It allows to create convenient methods With... that will add new value to the fields for
// that structure.
type VerifyOption func(*VerifyOptions)

// WithExternalID takes event identifier as a string, this is whatever the system wants to connect the proof
// with (e.g. email, phone number, incremental id, etc.)
func WithExternalID(identifier string) VerifyOption {
	return func(opts *VerifyOptions) {
		idHash := sha256.Sum256([]byte(identifier))
		opts.externalID = hex.EncodeToString(idHash[:])
	}
}

// WithAgeAbove adds new age check. It is an integer (e.g. 10, 18, 21) above which the person's
// age must be in proof.
func WithAgeAbove(age int) VerifyOption {
	return func(opts *VerifyOptions) {
		opts.age = time.Now().UTC().AddDate(-age, 0, 0)
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

// WithEventData takes some data for which proof was generated. This value format has to be validated before
// the signals validation because the kit checks ONLY the correspondence of these values.
func WithEventData(eventData interface{}) VerifyOption {
	return func(opts *VerifyOptions) {
		opts.eventData = eventData
	}
}

// WithEventID takes event identifier as a string that represents big number in a decimal format.
func WithEventID(identifier string) VerifyOption {
	return func(opts *VerifyOptions) {
		opts.eventID = identifier
	}
}

// WithRootVerifier takes an abstract verifier that should verify idStateRoot signal against identity tree
func WithRootVerifier(v IdentityRootVerifier) VerifyOption {
	return func(opts *VerifyOptions) {
		opts.rootVerifier = v
	}
}

// WithVerificationKeyFile takes a string that represents the name of the file
// with verification key. The file is read on NewPassportVerifier call. If you
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
// NOTE: WithIdentitiesCounter is tightly connected with WithIdentitiesCreationTimestampLimit.
// In case if only one of them is present, it is validated as usual, but if there are both
// options, one of them MAY be invalid. In a nutshell, they work as OR and not AND.
func WithIdentitiesCounter(maxIdentityCount int64) VerifyOption {
	return func(opts *VerifyOptions) {
		opts.maxIdentitiesCount = maxIdentityCount
	}
}

// WithIdentitiesCreationTimestampLimit takes the upper bound for timestamp when user might create
// identities.
//
// NOTE: WithIdentitiesCreationTimestampLimit is tightly connected with WithIdentitiesCounter.
// In case if only one of them is present, it is validated as usual, but if there are both
// options, one of them MAY be invalid. In a nutshell, they work as OR and not AND.
func WithIdentitiesCreationTimestampLimit(maxIdentityCreationTimestamp int64) VerifyOption {
	return func(opts *VerifyOptions) {
		opts.lastIdentityCreationTimestamp = time.Unix(maxIdentityCreationTimestamp, 0)
	}
}

// mergeOptions collects all parameters together and fills VerifyOptions struct
// with it, overwriting existing values
func mergeOptions(opts VerifyOptions, options ...VerifyOption) VerifyOptions {
	opts.maxIdentitiesCount = -1

	for _, opt := range options {
		opt(&opts)
	}

	if opts.rootVerifier == nil {
		opts.rootVerifier = identity.NewDisabledVerifier()
	}

	opts.allIdentitiesParamsSet = opts.maxIdentitiesCount == -1 || !opts.lastIdentityCreationTimestamp.IsZero()

	return opts
}

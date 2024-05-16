package zkverifier_kit

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	val "github.com/go-ozzo/ozzo-validation/v4"
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
	// eventDataRule - validation rule for EventData, where it's either an address or a string
	eventDataRule val.Rule
	// eventID - unique identifier associated with a specific event or interaction within
	// the protocol execution, may be used to keep track of various steps or actions, this
	// id is a string with a big integer in decimals format
	eventID string
	// rootVerifier - provider of identity root verification for IdStateRoot
	rootVerifier IdentityRootVerifier
	// verificationKeyFile - stores verification key for proofs
	verificationKeyFile string
	// maxIdentitiesCount - maximum amount of reissued identities that user can have
	maxIdentitiesCount int64
	// maxIdentityCreationTimestamp - the upper bound of timestamp when user could create identities
	maxIdentityCreationTimestamp time.Time
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

// WithEventData takes raw data for which the proof should be generated. This value format has to be validated before
// the signals validation because the kit checks ONLY the correspondence of these values.
func WithEventData(raw interface{}) VerifyOption {
	return func(opts *VerifyOptions) {
		opts.eventDataRule = matchesData(raw)
	}
}

// WithRarimoAddress takes decoded address that must be validated in proof. It
// requires to have same format that is in proof public signals (for example:
// bech32 address decoded to base256 without human-readable part)
//
// This should not be specified at the same time with WithEventData.
func WithRarimoAddress(address string) VerifyOption {
	return func(opts *VerifyOptions) {
		opts.eventDataRule = matchesAddress(address)
	}
}

// WithEventID takes event identifier as a string that represents big number in a decimal format.
func WithEventID(identifier string) VerifyOption {
	return func(opts *VerifyOptions) {
		opts.eventID = identifier
	}
}

// WithRootVerifier takes an abstract verifier that should verify IdStateRoot against the identity tree.
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

// mergeOptions collects all parameters together and fills VerifyOptions struct
// with it, overwriting existing values
func mergeOptions(opts VerifyOptions, options ...VerifyOption) VerifyOptions {
	opts.maxIdentitiesCount = -1
	opts.rootVerifier = identity.NewDisabledVerifier()

	for _, opt := range options {
		opt(&opts)
	}

	return opts
}

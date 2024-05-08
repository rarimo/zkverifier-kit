package zkverifier_kit

import (
	zkptypes "github.com/iden3/go-rapidsnark/types"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type proofType string

const PassportVerification proofType = "passport_proof"

var ErrUnknownProofType = errors.New("unknown proof type")

// Connector is an abstraction which collects all the methods to be implemented by each verifier
type Connector interface {
	SetExternalID(externalID string)
	VerifyExternalID(externalID string) error
	VerifyProof(proof zkptypes.ZKProof) error
}

// NewVerifier is a general constructor that will create a new verifier instance depending on
// proof type that was passed as argument, in its turn options have parameters that must be
// validated during proof verification, so they just transited to another constructor.
//
// See concrete Connector implementations to understand the handling of verificationKey arg.
func NewVerifier(pType proofType, verificationKey []byte, options ...VerifyOption) (Connector, error) {
	switch pType {
	case PassportVerification:
		return NewPassportVerifier(verificationKey, options...)
	default:
		return nil, errors.From(ErrUnknownProofType, logan.F{"type": string(pType)})
	}
}

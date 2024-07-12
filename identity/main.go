package identity

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

var (
	// ErrContractCall shows that the contract call on root validation was
	// unsuccessful, indicating an internal error
	ErrContractCall = errors.New("contract call failed")
	// ErrInvalidRoot shows that identity root verification flow completed
	// without internal errors, but the root itself is invalid
	ErrInvalidRoot = errors.New("invalid identity root")
)

// Verifier provides a method to verify user identity presence in the identity
// tree, stored in PoseidonSMT contract
type Verifier struct {
	caller   Caller
	timeout  time.Duration
	disabled bool
}

// Caller is an abstract contract caller, which verifiers identity root validity
type Caller interface {
	IsRootValid(opts *bind.CallOpts, root [32]byte) (bool, error)
}

func NewVerifier(caller Caller, timeout time.Duration) *Verifier {
	return &Verifier{
		caller:  caller,
		timeout: timeout,
	}
}

func NewDisabledVerifier() *Verifier {
	return &Verifier{disabled: true}
}

// VerifyRoot accepts an identity root from proof's pub signals as a big decimal
// integer, then calls the contract to check validity. It is recommended to
// have a special handling of ErrContractCall.
//
// If Verifier is disabled, nil is always returned.
func (v *Verifier) VerifyRoot(root string) error {
	if v.disabled {
		return nil
	}

	b, ok := new(big.Int).SetString(root, 10)
	if !ok {
		return ErrInvalidRoot
	}

	var provided [32]byte
	b.FillBytes(provided[:])

	ctx, cancel := context.WithTimeout(context.Background(), v.timeout)
	defer cancel()

	valid, err := v.caller.IsRootValid(&bind.CallOpts{Context: ctx}, provided)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrContractCall, err)
	}
	if !valid {
		return ErrInvalidRoot
	}

	return nil
}

// IsDisabled is useful when you want to have a different logic for disabled Verifier
func (v *Verifier) IsDisabled() bool {
	return v.disabled
}

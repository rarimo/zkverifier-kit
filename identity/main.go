package identity

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// ErrInvalidRoot indicates that identity root verification flow completed
// without internal errors, but the root itself is invalid
var ErrInvalidRoot = errors.New("invalid identity root")

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
// assert ErrInvalidRoot for a special internal errors handling.
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
	copy(provided[:], b.Bytes())

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(v.timeout))
	defer cancel()

	valid, err := v.caller.IsRootValid(&bind.CallOpts{Context: ctx}, provided)
	if err != nil {
		return fmt.Errorf("check root on contract: %w", err)
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

package voting

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"
)

// ErrInvalidRoot shows that identity root verification flow completed
// without internal errors, but the root itself is invalid
var ErrInvalidRoot = errors.New("invalid identity root")

// Verifier provides a method to verify user identity presence in the identity
// tree, stored in PoseidonSMT contract
type Verifier struct {
	validator RootValidator
	timeout   time.Duration
	disabled  bool
}

func NewVerifier(v RootValidator, timeout time.Duration) *Verifier {
	return &Verifier{
		validator: v,
		timeout:   timeout,
	}
}

func NewDisabledVerifier() *Verifier {
	return &Verifier{disabled: true}
}

// VerifyRoot accepts an identity root from proof's pub signals as a big decimal
// integer, then checks its presence among events from PoseidonSMT contract. It
// is recommended to have a special handling of ErrInvalidRoot.
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

	err := v.validator.ValidateRoot(ctx, provided)
	if err != nil {
		return fmt.Errorf("validate root: %w", err)
	}

	return nil
}

// IsDisabled is useful when you want to have a different logic for disabled Verifier
func (v *Verifier) IsDisabled() bool {
	return v.disabled
}

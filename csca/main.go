package csca

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// Verifier provides a method to verify CSCA tree root pub signal against the
// root stored on the Registration contract
type Verifier struct {
	caller   Caller
	timeout  time.Duration
	disabled bool
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

type Caller interface {
	IcaoMasterTreeMerkleRoot(*bind.CallOpts) ([32]byte, error)
}

// VerifyRoot accepts a root from proof's pub signals as a big decimal integer,
// calls the contract and compares two roots.
//
// If Verifier is disabled, nil is always returned.
func (v *Verifier) VerifyRoot(root string) error {
	if v.disabled {
		return nil
	}

	rootBig, ok := new(big.Int).SetString(root, 16)
	if !ok {
		return fmt.Errorf("invalid root passed: %s", root)
	}
	rootBytes := rootBig.Bytes()

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(v.timeout))
	defer cancel()

	raw, err := v.caller.IcaoMasterTreeMerkleRoot(&bind.CallOpts{Context: ctx})
	if err != nil {
		return fmt.Errorf("get root from contract: %w", err)
	}

	if !bytes.Equal(raw[:], rootBytes) {
		return fmt.Errorf("root mismatch: stored %x, provided %x", raw[:], rootBytes)
	}

	return nil
}

// IsDisabled is useful when you want to have a different logic for disabled Verifier
func (v *Verifier) IsDisabled() bool {
	return v.disabled
}

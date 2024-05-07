package csca

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

var ErrRootMismatch = fmt.Errorf("provided root does not match stored one")

// Verifier provides a method to verify CSCA tree root pub signal against the
// root stored on the Registration contract
type Verifier struct {
	caller   Caller
	timeout  time.Duration
	disabled bool
	cache    *cache
}

// Caller is an abstract contract caller, which retrieves current CSCA (ICAO) root
type Caller interface {
	IcaoMasterTreeMerkleRoot(*bind.CallOpts) ([32]byte, error)
}

type cache struct {
	root       []byte
	expiresAt  time.Time
	expiration time.Duration
}

func (c *cache) isExpired() bool {
	return c.expiresAt.Before(time.Now().UTC())
}

func (c *cache) update(root []byte) {
	c.root = root
	c.expiresAt = time.Now().UTC().Add(c.expiration)
}

func NewVerifier(caller Caller, timeout, cacheExpiration time.Duration) *Verifier {
	return &Verifier{
		caller:  caller,
		timeout: timeout,
		cache: &cache{
			expiration: cacheExpiration,
		},
	}
}

func NewDisabledVerifier() *Verifier {
	return &Verifier{disabled: true}
}

// VerifyRoot accepts a root from proof's pub signals as a big decimal integer,
// calls the contract and compares two roots.
//
// If Verifier is disabled, nil is always returned.
func (v *Verifier) VerifyRoot(root string) error {
	if v.disabled {
		return nil
	}

	b, ok := new(big.Int).SetString(root, 10)
	if !ok {
		return fmt.Errorf("invalid root passed: %s", root)
	}
	provided := b.Bytes()

	stored := v.cache.root
	if v.cache.isExpired() {
		ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(v.timeout))
		defer cancel()

		raw, err := v.caller.IcaoMasterTreeMerkleRoot(&bind.CallOpts{Context: ctx})
		if err != nil {
			return fmt.Errorf("get root from contract: %w", err)
		}

		stored = raw[:]
		v.cache.update(stored)
	}

	if !bytes.Equal(provided, stored) {
		return fmt.Errorf("%w: provided=%x, stored=%x", ErrRootMismatch, provided, stored)
	}

	return nil
}

// IsDisabled is useful when you want to have a different logic for disabled Verifier
func (v *Verifier) IsDisabled() bool {
	return v.disabled
}

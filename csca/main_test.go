package csca

import (
	"encoding/hex"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type mockCaller struct {
	root [32]byte
}

func (m *mockCaller) IcaoMasterTreeMerkleRoot(_ *bind.CallOpts) ([32]byte, error) {
	return m.root, nil
}

func TestVerifier_VerifyRoot(t *testing.T) {
	// converted from decimal providedRoot to hex storedRoot
	const (
		providedRoot = "16693841514009401027717517576091902513189966508499657428478303854796486502473"
		storedRoot   = "24e861243940eb879c33d91d1312bd0f7b44887342739eb210bdb30c01186849"
		expiration   = 3 * time.Second
	)

	testCases := []struct {
		name     string
		provided string // decimal
		stored   string // hex
		sleep    bool
		want     string
	}{
		{
			name:     "Should pass on the same root",
			provided: providedRoot,
			stored:   storedRoot,
			want:     "",
		},
		{
			name:     "Should fail on different root with ErrRootMismatch",
			provided: "166000000",
			stored:   storedRoot,
			want:     ErrRootMismatch.Error(),
		},
		{
			name:     "Should fail on invalid decimal",
			provided: "0x1234",
			stored:   storedRoot,
			want:     "invalid root passed: 0x123",
		},
		{
			name:     "Should pass on a different root when cache is fresh",
			provided: providedRoot,
			stored:   "ffffff",
			want:     "",
		},
		{
			name:     "Should fail on a different root when cache is expired",
			provided: providedRoot,
			stored:   "ffffff",
			sleep:    true,
			want:     ErrRootMismatch.Error(),
		},
	}

	caller := new(mockCaller)
	v := NewVerifier(caller, 0, expiration)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			caller.root = hexToBytes(tc.stored)

			if tc.sleep {
				v.cache.expiresAt = v.cache.expiresAt.Add(-expiration)
			}

			err := v.VerifyRoot(tc.provided)
			if err != nil {
				if tc.want == "" || !strings.Contains(err.Error(), tc.want) {
					t.Errorf("Verifier.VerifyRoot() = %q, want %q", err.Error(), tc.want)
				}
				return
			}

			if tc.want != "" {
				t.Errorf("Verifier.VerifyRoot() = <nil>, want %q", tc.want)
			}
		})
	}
}

func hexToBytes(h string) [32]byte {
	var b [32]byte
	bs, err := hex.DecodeString(h)
	if err != nil {
		panic(fmt.Errorf("failed to decode hex: %w", err))
	}

	copy(b[:], bs)
	return b
}

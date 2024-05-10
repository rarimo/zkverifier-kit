package identity

import (
	"testing"

	"github.com/rarimo/zkverifier-kit/internal/testutil"
	"github.com/stretchr/testify/assert"
)

func TestVerifier_VerifyRoot(t *testing.T) {
	// converted from decimal providedRoot to hex storedRoot
	const (
		providedRoot        = "16693841514009401027717517576091902513189966508499657428478303854796486502473"
		invalidProvidedRoot = "16693841000000000000000000006091902513189966508499657428478303854796486502473"
		storedRoot          = "24e861243940eb879c33d91d1312bd0f7b44887342739eb210bdb30c01186849"
	)

	testCases := []struct {
		name     string
		provided string // decimal
		stored   string // hex
		want     error
	}{
		{
			name:     "Should pass on the same root",
			provided: providedRoot,
			stored:   storedRoot,
			want:     nil,
		},
		{
			name:     "Should fail on different root",
			provided: invalidProvidedRoot,
			stored:   storedRoot,
			want:     ErrInvalidRoot,
		},
		{
			name:     "Should fail on invalid decimal",
			provided: "0x1234",
			stored:   storedRoot,
			want:     ErrInvalidRoot,
		},
	}

	caller := new(testutil.MockCaller)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := NewVerifier(caller.WithRoot(tc.stored), 0).VerifyRoot(tc.provided)
			assert.ErrorIs(t, err, tc.want)
		})
	}
}

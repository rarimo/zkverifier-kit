package zkverifier_kit

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"testing"

	zkptypes "github.com/iden3/go-rapidsnark/types"
	"github.com/rarimo/zkverifier-kit/identity"
	"github.com/rarimo/zkverifier-kit/internal/testutil"
	"github.com/stretchr/testify/assert"
)

// !!!NOTE: Tests will fail if ZKP is not generated at the same date, when these tests are run.
// This is because of the expiration_lower_bound check, that verifies that this signal is equal to
// the current date.

const (
	validAddress   = "rarimo1exzw7q2fytyrurkp5s7tm7ek720we9ejwujf2h"
	invalidAddress = "rarimo1nzmzvnr8yk98a9qxgkr0rrmmza7lhj90h9zycl"

	higherAge = 98
	lowerAge  = 13
	equalAge  = 18

	ukrCitizenship = "UKR"
	usaCitizenship = "USA"
	engCitizenship = "ENG"

	validEventID   = "304358862882731539112827930982999386691702727710421481944329166126417129570"
	invalidEventID = "AC42D1A986804618C7A793FBE814D9B31E47BE51E082806363DCA6958F3062"

	storedRoot = "1fd232b83b1927f2a8ede62ffe15c31d18782dd513e08f4aabeaf2e8e4c32417"

	maxTimestamp = math.MaxInt32
)

const verificationKeyFile = "example_verification_key.json"

var validProof = zkptypes.ZKProof{
	Proof: &zkptypes.ProofData{
		Protocol: "groth16",
		A: []string{
			"10580782106790373477261932143775321905443589586433651561536061428214589672484",
			"11698250021575150794852451179623994478616360899564591934730721390019157908788",
			"1",
		},
		B: [][]string{
			{
				"21705525629420011147308188725178223733293601519891005372105977395761857543648",
				"21430490393915822750736844419186297691953172223086304113869075235715014505478",
			},
			{
				"3610890604725359549886067582256338844864753769852189853964284902205861040749",
				"7808595281025880671232006108041019334015054900817419103023248582576073762625",
			},
			{
				"1",
				"0",
			},
		},
		C: []string{
			"14487193913320584434009947494902473354673034787917761246001827312658064548420",
			"13206968032646449669115920135803893331131897495922885759651807223610673459946",
			"1",
		},
	},
	PubSignals: []string{
		"7639957125598480790492529006924434106731566948760118579546114507674255247458",
		"0",
		"0",
		"0",
		"0",
		"0",
		"5589842",
		"0",
		"0",
		"304358862882731539112827930982999386691702727710421481944329166126417129570",
		"11318436481061661812577344400351359194387994145300108534310140806143276292370",
		"14393086243856018838405247242117964464658357003864077561407424514652280923159",
		"23073",
		"0",
		"1713436478",
		"0",
		"1",
		"52983525027888",
		"53009295159860",
		"55199728480820",
		"52983525027888",
		"0",
	},
}

// converted from EventData field in validProof.PubSignals
var (
	validEventData   = []byte{25, 6, 2, 14, 30, 0, 10, 9, 4, 11, 4, 3, 28, 3, 22, 1, 20, 16, 30, 11, 27, 30, 25, 22, 30, 10, 15, 14, 25, 5, 25, 18}
	invalidEventData = []byte{174}
)

var verificationKey []byte

func init() {
	var err error
	verificationKey, err = os.ReadFile(verificationKeyFile)
	if err != nil {
		panic(err)
	}
}

func TestNewPassportVerifier(t *testing.T) {
	testCases := []struct {
		name    string
		key     []byte
		keyFile string
		want    string
	}{
		{
			name: "Valid raw key",
			key:  verificationKey,
			want: "",
		},
		{
			name:    "Valid key from file",
			keyFile: verificationKeyFile,
			want:    "",
		},
		{
			name:    "Non-existent key file",
			keyFile: "nonexistent.json",
			want:    "failed to read verification key from file",
		},
		{
			name: "Neither key nor file are specified",
			want: ErrVerificationKeyRequired.Error(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			opt := make([]VerifyOption, 0, 1)
			if tc.keyFile != "" {
				opt = append(opt, WithVerificationKeyFile(tc.keyFile))
			}

			_, err := NewPassportVerifier(tc.key, opt...)
			if tc.want == "" {
				assert.NoError(t, err)
				return
			}

			assert.ErrorContains(t, err, tc.want)
		})
	}
}

func TestVerifyProof(t *testing.T) {
	var (
		defaultVerifier = identity.NewVerifier(new(testutil.MockCaller).WithRoot(storedRoot), 0)
		badVerifier     = identity.NewVerifier(new(testutil.MockCaller).WithRoot("ffffff"), 0)
		invalidKey      = bytes.Replace(verificationKey, []byte("1"), []byte("0"), -1)
	)

	testCases := []struct {
		name       string
		key        []byte
		initOpts   []VerifyOption
		verifyOpts []VerifyOption
		want       string
	}{
		{
			name: "Matching citizenship",
			initOpts: []VerifyOption{
				WithVerificationKeyFile(verificationKeyFile),
				WithProofSelectorValue("23073"),
				WithCitizenships(ukrCitizenship),
			},
			want: "",
		},
		{
			name: "Non-matching citizenship",
			initOpts: []VerifyOption{
				WithVerificationKeyFile(verificationKeyFile),
				WithProofSelectorValue("23073"),
				WithCitizenships(engCitizenship, usaCitizenship),
			},
			want: "pub_signals/citizenship: must be a valid value",
		},
		{
			name: "Valid address",
			initOpts: []VerifyOption{
				WithVerificationKeyFile(verificationKeyFile),
				WithProofSelectorValue("23073"),
				WithRarimoAddress(validAddress),
			},
			want: "",
		},
		{
			name: "Invalid address",
			initOpts: []VerifyOption{
				WithVerificationKeyFile(verificationKeyFile),
				WithProofSelectorValue("23073"),
				WithRarimoAddress(invalidAddress),
			},
			want: "pub_signals/event_data: must be a valid value",
		},
		{
			name: "Valid event data",
			initOpts: []VerifyOption{
				WithVerificationKeyFile(verificationKeyFile),
				WithProofSelectorValue("23073"),
				WithEventData(validEventData),
			},
			want: "",
		},
		{
			name: "Invalid event data",
			initOpts: []VerifyOption{
				WithVerificationKeyFile(verificationKeyFile),
				WithProofSelectorValue("23073"),
				WithEventData(invalidEventData),
			},
			want: "pub_signals/event_data: must be a valid value",
		},
		{
			name: "Lower age",
			initOpts: []VerifyOption{
				WithVerificationKeyFile(verificationKeyFile),
				WithProofSelectorValue("23073"),
				WithAgeAbove(lowerAge),
			}, verifyOpts: []VerifyOption{
				WithAgeAbove(lowerAge),
			},
			// Because proof is generated directly to the current_date - age (18 in our test case)
			want: "pub_signals/birth_date_upper_bound: dates are not equal",
		},
		{
			name: "Equal age",
			initOpts: []VerifyOption{
				WithVerificationKeyFile(verificationKeyFile),
				WithProofSelectorValue("23073"),
				WithAgeAbove(equalAge),
			},
			want: "",
		},
		{
			name: "Higher age",
			initOpts: []VerifyOption{
				WithVerificationKeyFile(verificationKeyFile),
				WithProofSelectorValue("23073"),
				WithAgeAbove(higherAge),
			},
			verifyOpts: []VerifyOption{
				WithAgeAbove(higherAge),
			},
			want: "pub_signals/birth_date_upper_bound: dates are not equal",
		},
		{
			name: "Valid event ID",
			initOpts: []VerifyOption{
				WithVerificationKeyFile(verificationKeyFile),
				WithProofSelectorValue("23073"),
				WithEventID(validEventID),
			},
			want: "",
		},
		{
			name: "Invalid event ID",
			initOpts: []VerifyOption{
				WithVerificationKeyFile(verificationKeyFile),
				WithProofSelectorValue("23073"),
				WithEventID(invalidEventID),
			},
			want: "pub_signals/event_id: must be a valid value",
		},
		{
			name: "Valid counter without timestamp",
			initOpts: []VerifyOption{
				WithVerificationKeyFile(verificationKeyFile),
				WithProofSelectorValue("23073"),
				WithIdentitiesCounter(999),
			},
			want: "",
		},
		{
			name: "Valid timestamp without counter",
			initOpts: []VerifyOption{
				WithVerificationKeyFile(verificationKeyFile),
				WithProofSelectorValue("23073"),
				WithIdentitiesCreationTimestampLimit(maxTimestamp),
			},
			want: "",
		},
		{
			name: "Valid counter with invalid timestamp",
			initOpts: []VerifyOption{
				WithVerificationKeyFile(verificationKeyFile),
				WithProofSelectorValue("23073"),
				WithIdentitiesCounter(999),
				WithIdentitiesCreationTimestampLimit(0),
			},
			want: "",
		},
		{
			name: "Valid timestamp with invalid counter",
			initOpts: []VerifyOption{
				WithVerificationKeyFile(verificationKeyFile),
				WithProofSelectorValue("23073"),
				WithIdentitiesCounter(0),
				WithIdentitiesCreationTimestampLimit(maxTimestamp),
			},
			want: "",
		},
		{
			name: "Invalid counter and timestamp",
			initOpts: []VerifyOption{
				WithVerificationKeyFile(verificationKeyFile),
				WithProofSelectorValue("23073"),
				WithIdentitiesCounter(0),
				WithIdentitiesCreationTimestampLimit(0),
			},
			verifyOpts: []VerifyOption{
				WithIdentitiesCounter(0),
				WithIdentitiesCreationTimestampLimit(1684839455),
			},
			want: "pub_signals/timestamp_upper_bound: must be no greater than 2023-05-23 13:57:35 +0300 EEST",
		},
		{
			name: "No options",
			initOpts: []VerifyOption{
				WithVerificationKeyFile(verificationKeyFile),
				WithProofSelectorValue("23073"),
			},
			want: "",
		},
		{
			name: "All valid options",
			initOpts: []VerifyOption{
				WithAgeAbove(equalAge),
				WithProofSelectorValue("23073"),
				WithCitizenships(ukrCitizenship),
				WithEventID(validEventID),
				WithIdentityVerifier(defaultVerifier),
				WithIdentitiesCounter(999),
				WithIdentitiesCreationTimestampLimit(maxTimestamp),
				WithVerificationKeyFile(verificationKeyFile),
			},
			verifyOpts: []VerifyOption{
				WithRarimoAddress(validAddress),
			},
			want: "",
		},
		{
			name: "Invalid identity verifier",
			initOpts: []VerifyOption{
				WithVerificationKeyFile(verificationKeyFile),
				WithIdentityVerifier(badVerifier),
				WithProofSelectorValue("23073"),
			},
			verifyOpts: []VerifyOption{
				WithVerificationKeyFile(verificationKeyFile),
				WithIdentityVerifier(badVerifier),
			},
			want: fmt.Sprintf("pub_signals/id_state_root: %s", identity.ErrInvalidRoot),
		},
		{
			name: "Invalid verification key",
			initOpts: []VerifyOption{
				WithProofSelectorValue("23073"),
			},
			key:  invalidKey,
			want: "groth16 verification failed",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			key := tc.key
			if len(key) == 0 {
				key = verificationKey
			}

			verifier, err := NewPassportVerifier(tc.key, tc.initOpts...)
			if err != nil {
				t.Fatal(err)
			}

			//err = verifier.VerifyProof(validProof, tc.verifyOpts...)
			verifier.opts = mergeOptions(false, verifier.opts, tc.verifyOpts...)
			err = verifier.validateBase(validProof)
			if tc.want == "" {
				assert.NoError(t, err)
				return
			}

			assert.ErrorContains(t, err, tc.want)
		})
	}
}

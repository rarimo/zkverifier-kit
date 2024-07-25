package zkverifier_kit

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"testing"

	zkptypes "github.com/iden3/go-rapidsnark/types"
	"github.com/rarimo/zkverifier-kit/root"
	"github.com/stretchr/testify/assert"
)

// !!!NOTE: Tests will fail if ZKP is not generated at the same date, when these tests are run.
// This is because of the expiration_lower_bound check, that verifies that this signal is equal to
// the current date.

const (
	higherAge = 98
	lowerAge  = 13
	equalAge  = 18

	ukrCitizenship = "UKR"
	usaCitizenship = "USA"
	engCitizenship = "ENG"

	validEventID   = "211985299740800702300256033401632392934377086534111448880928528431996790315"
	invalidEventID = "AC42D1A986804618C7A793FBE814D9B31E47BE51E082806363DCA6958F3062"

	storedRoot = "1fd232b83b1927f2a8ede62ffe15c31d18782dd513e08f4aabeaf2e8e4c32417"

	maxTimestamp = math.MaxInt32
)

const verificationKeyFile = "example_verification_key.json"

var validProof = zkptypes.ZKProof{
	Proof: &zkptypes.ProofData{
		A: []string{"12996016572939291630814335127078241511084725921569267221560959870881226842779",
			"761491852560139542885770222513604531448776143684161133298463831447152162597",
			"1",
		},
		B: [][]string{
			{
				"7599497178424536454186364289381353824249712189259655444764200186031301149772",
				"758126175915585817034687017843086369456608327321178876731878325331363448399",
			},
			{
				"6552077155141497820322593297915265785002325660175270629851305329247857575000",
				"9398527224370554058809232921797193562676853236790343293937085614424613225502",
			},
			{
				"1",
				"0",
			},
		},
		C: []string{
			"5520540944773640715610133783195440292316265697892695003821987613285146498487",
			"18240929522954357307751256221496990071805852196144718742156278323268552814849",
			"1",
		},
		Protocol: "groth16",
	},
	PubSignals: []string{
		"8558624556173674225153579865966359056844751077624402517288712777973816271870",
		"0",
		"0",
		"0",
		"0",
		"0",
		"5589842",
		"0",
		"0",
		"211985299740800702300256033401632392934377086534111448880928528431996790315",
		"1388928733714245704395968367156897439973029902189556257928944264186457955333",
		"16055556473534063237266173531819244891235292810222332599565643356062398286614",
		"23073",
		"0",
		"1719823039",
		"0",
		"3",
		"52983525027888",
		"53009295290417",
		"55199728611377",
		"52983525027888",
		"52983525027888",
	},
}

// converted from EventData field in validProof.PubSignals
var (
	validEventData   = []byte{3, 18, 27, 22, 5, 4, 24, 11, 14, 27, 18, 30, 20, 25, 23, 16, 5, 12, 16, 16, 26, 31, 20, 1, 28, 21, 20, 10, 13, 19, 28, 5}
	invalidEventData = []byte{4, 18, 27, 22, 5, 4, 24, 11, 14, 27, 18, 30, 20, 25, 23, 16, 5, 12, 16, 16, 26, 31, 20, 1, 28, 21, 20, 10, 13, 19, 28, 5}
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

			_, err := NewVerifier(tc.key, opt...)
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
		defaultVerifier = root.DisabledVerifier{}
		// TODO: add custom mock verifier
		badVerifier = root.DisabledVerifier{}
		invalidKey  = bytes.Replace(verificationKey, []byte("1"), []byte("0"), -1)
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
			want: "pub_signals/event_data: event data does not match.",
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
			// We need new proof for every day for pass this test
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
			// We need new proof for every day for pass this test, because of birthdat validation
			name: "All valid options",
			initOpts: []VerifyOption{
				WithProofType(GlobalPassport),
				WithAgeAbove(equalAge),
				WithProofSelectorValue("23073"),
				WithCitizenships(ukrCitizenship),
				WithEventID(validEventID),
				WithPassportRootVerifier(defaultVerifier),
				WithIdentitiesCounter(999),
				WithIdentitiesCreationTimestampLimit(maxTimestamp),
				WithVerificationKeyFile(verificationKeyFile),
			},
			want: "",
		},
		// TODO: disabled verifier don't return error
		// We need provide contract and rpc for bad verifier
		{
			name: "Invalid identity verifier",
			initOpts: []VerifyOption{
				WithVerificationKeyFile(verificationKeyFile),
				WithPassportRootVerifier(badVerifier),
				WithProofSelectorValue("23073"),
			},
			verifyOpts: []VerifyOption{
				WithVerificationKeyFile(verificationKeyFile),
				WithPassportRootVerifier(badVerifier),
			},
			want: fmt.Sprintf("pub_signals/id_state_root: %s", root.ErrInvalidRoot),
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

			tc.initOpts = append(tc.initOpts, withSkipExpirationCheck(true))
			verifier, err := NewVerifier(tc.key, tc.initOpts...)
			if err != nil {
				t.Fatal(err)
			}

			err = verifier.VerifyProof(validProof, tc.verifyOpts...)
			if tc.want == "" {
				assert.NoError(t, err)
				return
			}

			assert.ErrorContains(t, err, tc.want)
		})
	}
}

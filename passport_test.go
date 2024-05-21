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

const (
	validAddress   = "rarimo14c4vkfdv50gz9fqvjkcz4qjpm9z4sadmszucca"
	invalidAddress = "rarimo1nzmzvnr8yk98a9qxgkr0rrmmza7lhj90h9zycl"

	higherAge = 98
	lowerAge  = 13
	equalAge  = 18

	ukrCitizenship = "UKR"
	usaCitizenship = "USA"
	engCitizenship = "ENG"

	validEventID   = "304358862882731539112827930982999386691702727710421481944329166126417129570"
	invalidEventID = "AC42D1A986804618C7A793FBE814D9B31E47BE51E082806363DCA6958F3062"

	storedRoot = "1ca2515c70356a3b62e3a00e6f1fb0af4f5478a59de5d800d0efd8a74ec5467b"

	maxTimestamp = math.MaxInt32
)

const verificationKeyFile = "example_verification_key.json"

var validProof = zkptypes.ZKProof{
	Proof: &zkptypes.ProofData{
		Protocol: "groth16",
		A: []string{
			"18929392093012325347131052665407792211123081344400497915094341252476263438261",
			"8408679008273681595537212606093592786249494040078375479923024998257983071475",
			"1",
		},
		B: [][]string{
			{
				"15160749571539416435696026319722797986724507005425139887386580647177964433575",
				"418891762248400158424572797431315516884583570522212791159261025341957248366",
			},
			{
				"10121246100036896752109986908202239909550406172732565186372518849865546324107",
				"9655662684529702951082833477502777390806258408724141964907025445748892512786",
			},
			{
				"1",
				"0",
			},
		},
		C: []string{
			"6439412770130794205755637487074591576051810644474180957793569827360562352844",
			"6514662220472085416512552593928091396163871788691373442939864229679481297632",
			"1",
		},
	},
	PubSignals: []string{
		"13670197989959160947016892212488819355235823437209979068218084261720054582279",
		"52992115355956",
		"55216908480563",
		"0",
		"0",
		"0",
		"5589842",
		"0",
		"0",
		"304358862882731539112827930982999386691702727710421481944329166126417129570",
		"994318722035655867941976495378932234159094527419",
		"12951550518411690859840573908810811336996269038828192037883707959753719498363",
		"39",
		"15806704627620783043448169214838786348395809330456140685459045233186516590845",
	},
}

// converted from EventData field in validProof.PubSignals
var (
	validEventData   = []byte{174, 42, 203, 37, 172, 163, 208, 34, 164, 12, 149, 176, 42, 130, 65, 217, 69, 88, 117, 187}
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
		invalidKey      = bytes.Replace(verificationKey, []byte("1"), []byte("0"), 1)
	)

	testCases := []struct {
		name       string
		key        []byte
		initOpts   []VerifyOption
		verifyOpts []VerifyOption
		want       string
	}{
		{
			name:     "Matching citizenship",
			initOpts: []VerifyOption{WithCitizenships(ukrCitizenship)},
			want:     "",
		},
		{
			name:     "Non-matching citizenship",
			initOpts: []VerifyOption{WithCitizenships(engCitizenship, usaCitizenship)},
			want:     "pub_signals/citizenship: must be a valid value",
		},
		{
			name:     "Valid address",
			initOpts: []VerifyOption{WithRarimoAddress(validAddress)},
			want:     "",
		},
		{
			name:     "Invalid address",
			initOpts: []VerifyOption{WithRarimoAddress(invalidAddress)},
			want:     "pub_signals/event_data: event data does not match the address",
		},
		{
			name:     "Valid event data",
			initOpts: []VerifyOption{WithEventData(validEventData)},
			want:     "",
		},
		{
			name:     "Invalid event data",
			initOpts: []VerifyOption{WithEventData(invalidEventData)},
			want:     "pub_signals/event_data: must be a valid value",
		},
		{
			name:     "Lower age",
			initOpts: []VerifyOption{WithAgeAbove(lowerAge)},
			want:     "",
		},
		{
			name:     "Equal age",
			initOpts: []VerifyOption{WithAgeAbove(equalAge)},
			want:     "",
		},
		{
			name:     "Higher age",
			initOpts: []VerifyOption{WithAgeAbove(higherAge)},
			want:     "pub_signals/birth_date: date is too late",
		},
		{
			name:     "Valid event ID",
			initOpts: []VerifyOption{WithEventID(validEventID)},
			want:     "",
		},
		{
			name:     "Invalid event ID",
			initOpts: []VerifyOption{WithEventID(invalidEventID)},
			want:     "pub_signals/event_id: must be a valid value",
		},
		{
			name:     "Valid counter without timestamp",
			initOpts: []VerifyOption{WithIdentitiesCounter(999)},
			want:     "",
		},
		{
			name:     "Valid timestamp without counter",
			initOpts: []VerifyOption{WithIdentitiesCreationTimestampLimit(maxTimestamp)},
			want:     "",
		},
		{
			name:     "Valid counter with invalid timestamp",
			initOpts: []VerifyOption{WithIdentitiesCounter(999), WithIdentitiesCreationTimestampLimit(0)},
			want:     "",
		},
		{
			name:     "Valid timestamp with invalid counter",
			initOpts: []VerifyOption{WithIdentitiesCounter(0), WithIdentitiesCreationTimestampLimit(maxTimestamp)},
			want:     "",
		},
		{
			name:     "Invalid counter and timestamp",
			initOpts: []VerifyOption{WithIdentitiesCounter(0), WithIdentitiesCreationTimestampLimit(0)},
			want:     "pub_signals/timestamp_upper_bound: date is too late",
		},
		{
			name: "No options",
			want: "",
		},
		{
			name: "All valid options",
			initOpts: []VerifyOption{
				WithAgeAbove(equalAge),
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
				WithIdentityVerifier(badVerifier),
			},
			want: fmt.Sprintf("pub_signals/id_state_root: %s", identity.ErrInvalidRoot),
		},
		{
			name: "Invalid verification key",
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

			err = verifier.VerifyProof(validProof, tc.verifyOpts...)
			if tc.want == "" {
				assert.NoError(t, err)
				return
			}

			assert.ErrorContains(t, err, tc.want)
		})
	}
}

package zkverifier_kit

import (
	"os"
	"strings"
	"testing"
	"time"

	zkptypes "github.com/iden3/go-rapidsnark/types"
	"github.com/pkg/errors"
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

	hashedExternalID = "5f3d4868bb9c16dd83407eda63d5ce8f7ca39063df9eb9aef217e6c6ee9ffb20"

	verificationKeyFile = "example_verification_key.json"
)

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

var pubSignals = []string{
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
	"240515",
	"0",
	"1",
	"52983525027888",
	"52983525027888",
	"52983525027888",
	"52983525027888",
	"0",
}

var verificationKey []byte

func init() {
	var err error
	verificationKey, err = os.ReadFile(verificationKeyFile)
	if err != nil {
		panic(err)
	}
}

func TestWithCitizenship(t *testing.T) {
	verifier, err := NewPassportVerifier(verificationKey, WithCitizenships(ukrCitizenship))
	if err != nil {
		t.Fatal(errors.Wrap(err, "initiating new verifier failed"))
	}

	if err = verifier.VerifyProof(validProof); err != nil {
		t.Fatal(errors.Wrap(err, "verifying proof"))
	}
}

func TestWithCitizenshipFail(t *testing.T) {
	verifier, err := NewPassportVerifier(verificationKey, WithCitizenships(usaCitizenship, engCitizenship))
	if err != nil {
		t.Fatal(errors.Wrap(err, "initiating new verifier failed"))
	}

	if err = verifier.VerifyProof(validProof); err != nil {
		if !assert.Equal(t, err.Error(), "pub_signals/citizenship: must be a valid value.") {
			t.Fatal(errors.Wrap(err, "verifying proof"))
		}
	}
}

func TestWithRarimoAddress(t *testing.T) {
	verifier, err := NewPassportVerifier(verificationKey, WithEventData(validAddress))
	if err != nil {
		t.Fatal(errors.Wrap(err, "initiating new verifier failed"))
	}

	if err = verifier.VerifyProof(validProof); err != nil {
		t.Fatal(errors.Wrap(err, "verifying proof"))
	}
}

func TestWithRarimoAddressFail(t *testing.T) {
	verifier, err := NewPassportVerifier(verificationKey, WithEventData(invalidAddress))
	if err != nil {
		t.Fatal(errors.Wrap(err, "initiating new verifier failed"))
	}

	if err = verifier.VerifyProof(validProof); err != nil {
		if !assert.Equal(t, err.Error(), "pub_signals/event_data: event data does not match the address.") {
			t.Fatal(errors.Wrap(err, "verifying proof"))
		}
	}
}

func TestWithAgeLower(t *testing.T) {
	verifier, err := NewPassportVerifier(verificationKey, WithAgeAbove(lowerAge))
	if err != nil {
		t.Fatal(errors.Wrap(err, "initiating new verifier failed"))
	}

	if err = verifier.VerifyProof(validProof); err != nil {
		t.Fatal(errors.Wrap(err, "verifying proof"))
	}
}

func TestWithAgeEqual(t *testing.T) {
	verifier, err := NewPassportVerifier(verificationKey, WithAgeAbove(equalAge))
	if err != nil {
		t.Fatal(errors.Wrap(err, "initiating new verifier failed"))
	}

	if err = verifier.VerifyProof(validProof); err != nil {
		t.Fatal(errors.Wrap(err, "verifying proof"))
	}
}

func TestWithAgeHigher(t *testing.T) {
	verifier, err := NewPassportVerifier(verificationKey, WithAgeAbove(higherAge))
	if err != nil {
		t.Fatal(errors.Wrap(err, "initiating new verifier failed"))
	}

	if err = verifier.VerifyProof(validProof); err != nil {
		if !assert.Equal(t, err.Error(), "pub_signals/birth_date: date is too late.") {
			t.Fatal(errors.Wrap(err, "verifying proof"))
		}
	}
}

func TestWithEventID(t *testing.T) {
	verifier, err := NewPassportVerifier(verificationKey, WithEventID(validEventID))
	if err != nil {
		t.Fatal(errors.Wrap(err, "initiating new verifier failed"))
	}

	if err = verifier.VerifyProof(validProof); err != nil {
		t.Fatal(errors.Wrap(err, "verifying proof"))
	}
}

func TestWithInvalidEventID(t *testing.T) {
	verifier, err := NewPassportVerifier(verificationKey, WithEventID(invalidEventID))
	if err != nil {
		t.Fatal(errors.Wrap(err, "initiating new verifier failed"))
	}

	if err = verifier.VerifyProof(validProof); err != nil {
		if !assert.Equal(t, err.Error(), "pub_signals/event_id: must be a valid value.") {
			t.Fatal(errors.Wrap(err, "verifying proof"))
		}
	}
}

func TestWithExternalID(t *testing.T) {
	verifier, err := NewPassportVerifier(verificationKey, WithExternalID(validAddress))
	if err != nil {
		t.Fatal(errors.Wrap(err, "initiating new verifier failed"))
	}

	if err = verifier.VerifyProof(validProof, WithExternalID(hashedExternalID)); err != nil {
		t.Fatal(errors.Wrap(err, "verifying proof"))
	}
}

func TestIdentitiesParams(t *testing.T) {
	verifier, err := NewPassportVerifier(verificationKey,
		WithIdentitiesCounter(0),
		WithIdentitiesCreationTimestampLimit(time.Now().Unix()),
	)
	if err != nil {
		t.Fatal(errors.Wrap(err, "initiating new verifier failed"))
	}

	err = verifier.validateIdentitiesInputs(pubSignals)
	if err != nil {
		t.Fatal(errors.Wrap(err, "validating inputs"))
	}
}

func TestInvalidCounterIdentityParams(t *testing.T) {
	inputsCopy := pubSignals
	inputsCopy[16] = "100"

	verifier, err := NewPassportVerifier(verificationKey,
		WithIdentitiesCounter(3),
		WithIdentitiesCreationTimestampLimit(1000),
	)
	if err != nil {
		t.Fatal(errors.Wrap(err, "initiating new verifier failed"))
	}

	err = verifier.validateIdentitiesInputs(inputsCopy)
	if err != nil {
		t.Fatal(errors.Wrap(err, "validating inputs"))
	}
}

func TestInvalidIdentitiesParams(t *testing.T) {
	inputsCopy := pubSignals
	inputsCopy[14] = "5000"
	inputsCopy[16] = "100"

	verifier, err := NewPassportVerifier(verificationKey,
		WithIdentitiesCounter(3),
		WithIdentitiesCreationTimestampLimit(1000),
	)
	if err != nil {
		t.Fatal(errors.Wrap(err, "initiating new verifier failed"))
	}

	if err = verifier.validateIdentitiesInputs(inputsCopy); err != nil {
		if !assert.Equal(t, err.Error(), "pub_signals/identity_counter_upperbound: must be no greater than 3.\npub_signals/timestamp_upperbound: must be no greater than 1000.") {
			t.Fatal(errors.Wrap(err, "validating inputs"))
		}
	}
}

func TestWithInvalidExternalID(t *testing.T) {
	verifier, err := NewPassportVerifier(verificationKey)
	if err != nil {
		t.Fatal(errors.Wrap(err, "initiating new verifier failed"))
	}

	if err = verifier.VerifyProof(validProof, WithExternalID(hashedExternalID)); err != nil {
		if !assert.Equal(t, err.Error(), "external_id: must be a valid value.") {
			t.Fatal(errors.Wrap(err, "verifying proof"))
		}
	}
}

func TestWithManyOptions(t *testing.T) {
	rootVerifier := identity.NewVerifier(new(testutil.MockCaller).WithRoot(storedRoot), 0)

	verifier, err := NewPassportVerifier(
		verificationKey,
		WithAgeAbove(equalAge),
		WithCitizenships(ukrCitizenship),
		WithEventID(validEventID),
		WithRootVerifier(rootVerifier),
		WithVerificationKeyFile(verificationKeyFile),
	)
	if err != nil {
		t.Fatal(errors.Wrap(err, "initiating new verifier failed"))
	}

	err = verifier.VerifyProof(validProof, WithEventData(validAddress))
	if err != nil {
		t.Fatal(errors.Wrap(err, "verifying proof"))
	}
}

func TestWithManyOptionsFail(t *testing.T) {
	rootVerifier := identity.NewVerifier(new(testutil.MockCaller).WithRoot("ffffff"), 0)

	verifier, err := NewPassportVerifier(
		nil,
		WithAgeAbove(higherAge),
		WithCitizenships(usaCitizenship),
		WithEventID(invalidEventID),
		WithRootVerifier(rootVerifier),
		WithVerificationKeyFile(verificationKeyFile),
	)
	if err != nil {
		t.Fatal(errors.Wrap(err, "initiating new verifier failed"))
	}

	if err = verifier.VerifyProof(validProof, WithEventData(invalidAddress)); err != nil {
		if !assert.Equal(t, err.Error(), "pub_signals/birth_date: date is too late; pub_signals/citizenship: must be a valid value; pub_signals/event_data: event data does not match the address; pub_signals/event_id: must be a valid value; pub_signals/id_state_hash: invalid identity root.") {
			t.Fatal(errors.Wrap(err, "verifying proof"))
		}
	}
}

func TestInvalidVerificationKey(t *testing.T) {
	_, err := NewPassportVerifier(nil)
	if err == nil || err.Error() != ErrVerificationKeyRequired.Error() {
		t.Errorf("NewPassportVerifier(nil) = %v, expected %s", err, ErrVerificationKeyRequired.Error())
	}

	_, err = NewPassportVerifier(nil, WithVerificationKeyFile("nonexistent"))
	if err == nil || !strings.Contains(err.Error(), "failed to read verification key from file") {
		t.Errorf("NewPassportVerifier(nil) = %v, expected %s", err, ErrVerificationKeyRequired.Error())
	}
}

func TestInvalidProofType(t *testing.T) {
	if _, err := NewVerifier("unknown_type", verificationKey); err != nil {
		if !assert.Error(t, ErrUnknownProofType, err) {
			t.Fatal(errors.Wrap(err, "initiating new verifier failed"))
		}
	}
}

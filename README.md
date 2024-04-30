# ZKVerifier Kit

## Introduction
This repository is an SDK for connecting external services with Rarimo passport
verification system that may be used in back-end services. Its main purpose is 
providing a convenient methods that can be used in any system without deep
introduction in the Rarimo system structure.

## Usage

General use case for usage looks like:
```go
    verifier, err := NewVerifier(
		PassportVerification, 
		WithEventID("304358862882731539112827930982999386691702727710421481944329166126417129570"),
		WithAgeAbove(18),
		WithCitizenships("UKR"),
    )
    if err != nil {
        return errors.Wrap(err, "failed to create new verifier")
    }

    if err = verifier.VerifyProof(proof); err != nil {
        return errors.Wrap(err, "failed to verify proof")
    }
```
Provided example will create a new instance of verifier that takes options of the lowes
available age, citizenship and event id that have to be validated before proof verification. 

Moreover, you may set and verify external identifier to connect proof with. There is an ability
to set this id during initialization and change later or just set when verifier has been already
created:
```go
    verifier, err := NewVerifier(
        WithExternalID("550e8400-e29b-41d4-a716-446655440000"),
    )
    if err != nil {
        return errors.Wrap(err, "failed to create new verifier")
    }

    verifier.SetExternalID(identifier) // this will override previous declaration of externalID
	
    externalIDHash := sha256.Sum256([]byte(identifier))
    if err = verifier.VerifyExternalID(hex.EncodeToString(externalIDHash[:])); err != nil {
        return errors.Wrap(err, "failed to verify external identifier")
    }
```
It worth to notice that to verify external identifier you have to take SHA256 hash from the raw
value that was passed in the `WithExternalID(...)` or `SetExternalID(...)` 


More usage example can be found in [verifier tests](passport_test.go).


Proof can be gained from the front-end apps or related Rarimo mobile applications. In general,
it has such format:

```go
    type ZKProof struct {
        Proof struct {
            A        []string   `json:"pi_a"`
            B        [][]string `json:"pi_b"`
            C        []string   `json:"pi_c"`
            Protocol string     `json:"protocol"`
        } `json:"proof"`
        PubSignals []string `json:"pub_signals"`
    }
```
In our systems mostly used ZKProof type from the [iden3 package.](https://github.com/iden3/go-rapidsnark)

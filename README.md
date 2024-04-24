# Integration SDK

## Introduction
This repository is an SDK for connecting external services with Rarimo passport
verification system that may be used in back-end services. Its main purpose is 
providing a convenient methods that can be used in any system without deep
introduction in the Rarimo system structure.

## Usage

General use case for usage looks like:
```go
    verifier, err := NewVerifier(PassportVerification, WithAgeAbove(18), WithCitizenships("UKR"))
    if err != nil {
        return errors.Wrap(err, "failed to create new verifier")
    }
    
    if err = verifier.VerifyProof(proof); err != nil {
        return errors.Wrap(err, "failed to verify proof")
    }
```
Provided example will create a new instance of verifier that takes options of the lowes
available age and citizenship that have to be validated before proof verification. More 
usage example can be found in [verifier tests](passport_verifier_test.go).


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
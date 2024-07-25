# ZKVerifier Kit

## Introduction

This repository is an SDK for connecting external services with Rarimo passport
verification and voting systems that may be used in back-end services. Its main purpose is 
providing a convenient methods that can be used in any system without deep
introduction in the Rarimo system structure.

## Usage

General usage looks like this:

```go
package main

import (
	kit "github.com/rarimo/zkverifier-kit"
	"github.com/rarimo/zkverifier-kit/root"
)

func main() {
	passportVerifier := root.NewPoseidonSMTVerifier(rpcURL, contractAddress, reqTimeout)
	
	v, err := kit.NewVerifier(
		nil,
		kit.WithProofType(kit.GeneralPassport),
		kit.WithVerificationKeyFile("key.json"),
		kit.WithEventID("304358862882731539112827930982999386691702727710421481944329166126417129570"),
		kit.WithAgeAbove(18),
		kit.WithCitizenships("UKR"),
		kit.WithPassportRootVerifier(passportVerifier),
		kit.WithIdentitiesCounter(0),
		kit.WithIdentitiesCreationTimestampLimit(1847321000),
	)
	if err != nil {
		// ...
	}
	// data is an abstract event data that you expect to be in proof
	err = v.VerifyProof(proof, kit.WithEventData(data))
	if err != nil {
		// ...
	}
}
```

```go
package main

import (
	kit "github.com/rarimo/zkverifier-kit"
	"github.com/rarimo/zkverifier-kit/root"
)

func main() {
	pollVerifier := root.NewProposalSMTVerifier(rpcURL, reqTimeout)
	
	v, err := kit.NewVerifier(
		nil,
		kit.WithProofType(kit.PollParticipation),
		kit.WithVerificationKeyFile("key.json"),
		kit.WithEventID("304358862882731539112827930982999386691702727710421481944329166126417129570"),
	)
	if err != nil {
		// ...
	}
	// data is an abstract event data that you expect to be in proof
	err = v.VerifyProof(proof,
		kit.WithPollParticipationEventID(pollEventID) ,
		kit.WithPollRootVerifier(v.WithContract(pollContract)))
	if err != nil {
		// ...
	}
}
```
Let's break this down.

### Configurable root verifier

Firstly, you instantiate identity root verifier, which will verify the
`IdStateRoot` public signal with contract call. You can refer to our
generated contract bindings in [poseidonsmt](internal/poseidonsmt) and [proposalsmt](internal/proposalsmt) package.
However, maybe, you would like to create the verifier from config map.

Here is configuration sample that you should have in `config.yaml` of your app:
```yaml
verifier:
  allowed_age: 18
  allowed_identity_timestamp: 1715698750
poseidonsmt_root_verifier:
  rpc: evm_rpc_url
  contract: poseidon_smt_contract_address
  request_timeout: 10s
proposalsmt_root_verifier:
  rpc: evm_rpc_url
  request_timeout: 10s
```

You can get values with [gitlab.com/distributed_lab/kit/kv](https://gitlab.com/distributed_lab/kit/-/tree/master/kv?ref_type=heads) package.
Then just create the verifier from config:
```go
    getter := kv.MustFromEnv()
    config := root.NewVerifierProvider(getter, root.PoseidonSMT)
	rv := config.ProvideVerifier()
```

### Custom verification key

If you specify `WithVerificationKeyPath`, the app will try to open the file and
convert its contents to bytes. Check out `example_verification_key.json`. You
don't have to get the key this way: just convert your key to bytes and pass it
directly without the mentioned option:
```go
v, err := kit.NewVerifier(
	kit.PassportVerification,
	keyBytes,
	options...
)
```

### Notes about options

Each option adds new validation rule to the proof, except `WithVerificaitonKeyFile`. Most of the options can be combined, but here is what you should consider:
- If you pass non-nil verification key, don't use `WithVerificationKeyFile`
- Don't use `WithEventData` together with `WithRarimoAddress`, because the address check is basically the data check with extra validation
- It is recommended to use `WithIdentitiesCounter` and `WithIdentitiesCreationTimestampLimit` together, because they imply a shared business logic of protection against double-eligibility.

You have two ways of providing options: globally (`NewVerifier`, `NewPassportVerifier`) and locally (`VerifyProof`). The latter override the former.

More usage examples can be found in [verifier tests](passport_test.go).

## Proof format

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
In our systems mostly used ZKProof type is the one from [iden3 package](https://github.com/iden3/go-rapidsnark).

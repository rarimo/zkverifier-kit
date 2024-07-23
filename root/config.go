package root

import (
	"fmt"
	"time"

	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

const baseTimeout = 5 * time.Second

// VerifierProvider provides a Verifier based on the given VerifierType from
// config map. You must specify the name equal to VerifierType in map: this
// allows to have multiple verifiers in the same app. For custom name or logic
// write your own config map handler.
//
// Specifying "disabled: true" allows to skip other map fields.
type VerifierProvider interface {
	ProvideVerifier(VerifierType) Verifier
}

type config struct {
	once   comfig.Once
	getter kv.Getter
}

func NewVerifierProvider(getter kv.Getter) VerifierProvider {
	return &config{getter: getter}
}

func (c *config) ProvideVerifier(typ VerifierType) Verifier {
	return c.once.Do(func() interface{} {
		var disabled struct {
			Disabled bool `fig:"disabled"`
		}

		err := figure.Out(&disabled).
			From(kv.MustGetStringMap(c.getter, string(typ))).
			Please()
		if err != nil {
			panic(fmt.Errorf("failed to figure out %s disabled field: %w", typ, err))
		}
		if disabled.Disabled {
			return DisabledVerifier{}
		}

		var cfg struct {
			RPC            string        `fig:"rpc,required"`
			Contract       string        `fig:"contract,required"`
			RequestTimeout time.Duration `fig:"request_timeout"`
		}

		err = figure.Out(&cfg).
			With(figure.EthereumHooks).
			From(kv.MustGetStringMap(c.getter, "root_verifier")).
			Please()
		if err != nil {
			panic(fmt.Errorf("failed to figure out root_verifier: %w", err))
		}

		if cfg.RequestTimeout == 0 {
			cfg.RequestTimeout = baseTimeout
		}

		var v Verifier
		switch typ {
		case PoseidonSMT:
			v, err = NewPoseidonSMTVerifier(cfg.RPC, cfg.Contract, cfg.RequestTimeout)
		case ProposalSMT:
			v, err = NewProposalSMTVerifier(cfg.RPC, cfg.Contract, cfg.RequestTimeout)
		default:
			panic(fmt.Errorf("unsupported verifier type: %s", typ))
		}

		if err != nil {
			panic(fmt.Errorf("failed to create %s verifier: %w", typ, err))
		}

		return v
	}).(Verifier)
}

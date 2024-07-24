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
// config map.
//
// Specifying "disabled: true" in config allows to skip other map fields.
type VerifierProvider interface {
	ProvideVerifier() Verifier
}

type config struct {
	once   comfig.Once
	getter kv.Getter
	typ    VerifierType
}

// NewVerifierProvider creates a new provider with given VerifierType. You must
// specify the name equal to VerifierType in map: this allows to have multiple
// verifiers in the same app. For custom name or logic write your own config map
// handler.
func NewVerifierProvider(getter kv.Getter, typ VerifierType) VerifierProvider {
	switch typ {
	case PoseidonSMT, ProposalSMT:
	default:
		panic(fmt.Errorf("unsupported verifier type: %s", typ))
	}
	return &config{getter: getter, typ: typ}
}

func (c *config) ProvideVerifier() Verifier {
	return c.once.Do(func() interface{} {
		var disabled struct {
			Disabled bool `fig:"disabled"`
		}

		err := figure.Out(&disabled).
			From(kv.MustGetStringMap(c.getter, string(c.typ))).
			Please()
		if err != nil {
			panic(fmt.Errorf("failed to figure out %s disabled field: %w", c.typ, err))
		}
		if disabled.Disabled {
			return DisabledVerifier{}
		}

		var cfg struct {
			RPC            string        `fig:"rpc,required"`
			Contract       string        `fig:"contract"`
			RequestTimeout time.Duration `fig:"request_timeout"`
		}

		err = figure.Out(&cfg).
			With(figure.EthereumHooks).
			From(kv.MustGetStringMap(c.getter, string(c.typ))).
			Please()
		if err != nil {
			panic(fmt.Errorf("failed to figure out %s: %w", c.typ, err))
		}

		if cfg.RequestTimeout == 0 {
			cfg.RequestTimeout = baseTimeout
		}

		var v Verifier
		switch c.typ {
		case PoseidonSMT:
			v, err = NewPoseidonSMTVerifier(cfg.RPC, cfg.Contract, cfg.RequestTimeout)
		case ProposalSMT:
			v = NewProposalSMTVerifier(cfg.RPC, cfg.RequestTimeout).WithContract(cfg.Contract)
		}

		if err != nil {
			panic(fmt.Errorf("failed to create %s verifier: %w", c.typ, err))
		}

		return v
	}).(Verifier)
}

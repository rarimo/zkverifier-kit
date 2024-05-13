package identity

import (
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rarimo/zkverifier-kit/internal/poseidonsmt"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

const baseTimeout = 5 * time.Second

type VerifierProvider struct {
	once   *comfig.Once
	getter kv.Getter
}

func NewVerifierProvider(getter kv.Getter) VerifierProvider {
	return VerifierProvider{
		getter: getter,
		once:   new(comfig.Once),
	}
}

func (c VerifierProvider) ProvideVerifier() *Verifier {
	return c.once.Do(func() interface{} {
		var disabled struct {
			Disabled bool `fig:"disabled"`
		}

		err := figure.Out(&disabled).
			From(kv.MustGetStringMap(c.getter, "root_verifier")).
			Please()
		if err != nil {
			panic(fmt.Errorf("failed to figure out root_verifier disabled field: %s", err))
		}
		if disabled.Disabled {
			return NewDisabledVerifier()
		}

		var cfg struct {
			RPC            string         `fig:"rpc,required"`
			Contract       common.Address `fig:"contract,required"`
			RequestTimeout time.Duration  `fig:"request_timeout"`
		}

		err = figure.Out(&cfg).
			With(figure.EthereumHooks).
			From(kv.MustGetStringMap(c.getter, "root_verifier")).
			Please()
		if err != nil {
			panic(fmt.Errorf("failed to figure out root_verifier: %s", err))
		}

		if cfg.RequestTimeout == 0 {
			cfg.RequestTimeout = baseTimeout
		}

		cli, err := ethclient.Dial(cfg.RPC)
		if err != nil {
			panic(fmt.Errorf("failed to connect to rpc: %w", err))
		}

		caller, err := poseidonsmt.NewPoseidonSMTCaller(cfg.Contract, cli)
		if err != nil {
			panic(fmt.Errorf("failed to bind registration contract caller: %w", err))
		}

		return NewVerifier(caller, cfg.RequestTimeout)
	}).(*Verifier)
}

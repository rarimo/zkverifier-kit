package csca

import (
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rarimo/zkverifier-kit/internal/registration"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

const (
	baseTimeout    = 5 * time.Second
	baseExpiration = 10 * time.Second
)

type Config interface {
	RootVerifier() *Verifier
}

type config struct {
	once   comfig.Once
	getter kv.Getter
}

func NewConfig(getter kv.Getter) Config {
	return &config{getter: getter}
}

func (c *config) RootVerifier() *Verifier {
	return c.once.Do(func() interface{} {
		var disabled struct {
			Disabled bool `fig:"disabled"`
		}

		err := figure.Out(&disabled).
			From(kv.MustGetStringMap(c.getter, "csca_verifier")).
			Please()
		if err != nil {
			panic(fmt.Errorf("failed to figure out csca_verifier disabled field: %s", err))
		}
		if disabled.Disabled {
			return NewDisabledVerifier()
		}

		var cfg struct {
			RPC             string         `fig:"rpc,required"`
			Contract        common.Address `fig:"contract,required"`
			RequestTimeout  time.Duration  `fig:"request_timeout"`
			CacheExpiration time.Duration  `fig:"cache_expiration"`
		}

		err = figure.Out(&cfg).
			With(figure.EthereumHooks).
			From(kv.MustGetStringMap(c.getter, "csca_verifier")).
			Please()
		if err != nil {
			panic(fmt.Errorf("failed to figure out csca_verifier: %s", err))
		}

		if cfg.RequestTimeout == 0 {
			cfg.RequestTimeout = baseTimeout
		}
		if cfg.CacheExpiration == 0 {
			cfg.RequestTimeout = baseExpiration
		}

		cli, err := ethclient.Dial(cfg.RPC)
		if err != nil {
			panic(fmt.Errorf("failed to connect to rpc: %w", err))
		}

		caller, err := registration.NewRegistrationCaller(cfg.Contract, cli)
		if err != nil {
			panic(fmt.Errorf("failed to bind registration contract caller: %w", err))
		}

		return NewVerifier(caller, cfg.RequestTimeout, cfg.CacheExpiration)
	}).(*Verifier)
}

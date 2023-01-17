package config

import (
	"github.com/Swapica/relayer-svc/internal/tx"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (c *config) Chains() tx.Chains {
	return c.chains.Do(func() interface{} {
		var cfg struct {
			tx.Chains `fig:"list,required"`
		}

		err := figure.Out(&cfg).
			With(figure.EthereumHooks).
			From(kv.MustGetStringMap(c.getter, "chains")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out chains"))
		}

		return cfg.Chains
	}).(tx.Chains)
}

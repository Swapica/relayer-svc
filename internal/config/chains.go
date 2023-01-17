package config

import (
	"github.com/Swapica/relayer-svc/internal/tx"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (c *config) Chains() tx.Chains {
	return c.chains.Do(func() interface{} {
		const errFigOut = "failed to figure out chains"
		var cfg struct {
			tx.Chains `fig:"list,required"`
		}

		err := figure.Out(&cfg).
			With(figure.EthereumHooks).
			From(kv.MustGetStringMap(c.getter, "chains")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, errFigOut))
		}
		if len(cfg.Chains) == 0 {
			panic(errFigOut + ": at least one chain must be present")
		}

		ids := make([]int64, len(cfg.Chains))
		names := make([]string, len(cfg.Chains))
		for i, ch := range cfg.Chains {
			ids[i] = ch.ID
			names[i] = ch.Name
		}
		if !isSet(ids) || !isSet(names) {
			panic(errFigOut + ": same chain IDs or names must not be present")
		}

		return cfg.Chains
	}).(tx.Chains)
}

func isSet[T comparable](arr []T) bool {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				return false
			}
		}
	}
	return true
}

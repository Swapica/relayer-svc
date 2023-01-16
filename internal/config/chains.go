package config

import (
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type Chain struct {
	ID       int64          `fig:"id,required"`
	Contract common.Address `fig:"contract,required"`
	RPC      string         `fig:"rpc,required"`
}

type Chains []Chain

func (c *config) Chains() Chains {
	return c.chains.Do(func() interface{} {
		var cfg struct {
			Chains `fig:"list,required"`
		}

		err := figure.Out(&cfg).
			With(figure.EthereumHooks).
			From(kv.MustGetStringMap(c.getter, "chains")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out chains"))
		}

		return cfg.Chains
	}).(Chains)
}

func (chains Chains) Get(id int64) *Chain {
	for _, c := range chains {
		// be careful with loop vars, probably better to use index
		if c.ID == id {
			return &c
		}
	}
	return nil
}

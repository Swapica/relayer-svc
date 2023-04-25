package tx

import "github.com/ethereum/go-ethereum/common"

type Chainer interface {
	Chains() Chains
}

type Chain struct {
	Name     string         `fig:"name,required"`
	Contract common.Address `fig:"contract,required"`
	RPC      string         `fig:"rpc,required"`
	ID       int64          `fig:"chain_id,required"`
}

type Chains []Chain

func (ch Chains) Get(name string) *Chain {
	for _, c := range ch {
		if c.Name == name {
			return &c
		}
	}
	return nil
}

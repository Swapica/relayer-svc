package tx

import "github.com/ethereum/go-ethereum/common"

type Chainer interface {
	Chains() Chains
}

type Chain struct {
	ID       int64          `fig:"id,required"`
	Contract common.Address `fig:"contract,required"`
	RPC      string         `fig:"rpc,required"`
}

type Chains []Chain

func (ch Chains) Get(id int64) *Chain {
	for _, c := range ch {
		if c.ID == id {
			return &c
		}
	}
	return nil
}

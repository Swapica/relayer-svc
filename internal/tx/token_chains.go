package tx

import (
	"github.com/Swapica/relayer-svc/resources"
	"strings"
)

type TokenChainer interface {
	TokenChains() TokenChains
}

type TokenChains []resources.TokenChain

func (tcs TokenChains) Get(tokenAddr, chain string) *resources.TokenChain {
	for _, t := range tcs {
		if strings.ToLower(t.Attributes.ContractAddress) == strings.ToLower(tokenAddr) &&
			t.Attributes.ChainId == chain {
			return &t
		}
	}
	return nil
}

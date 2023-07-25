package config

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/Swapica/relayer-svc/internal/tx"
	"github.com/Swapica/relayer-svc/resources"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (c *config) TokenChains() tx.TokenChains {
	return c.chains.Do(func() interface{} {
		var cfg struct {
			*url.URL `fig:"url,required"`
		}

		err := figure.Out(&cfg).From(kv.MustGetStringMap(c.getter, "token_chains")).Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out chains endpoint"))
		}

		resp, err := http.Get(cfg.URL.String())
		if err != nil {
			panic(errors.Wrap(err, "failed to fetch chain list"))
		}

		var tokenChains resources.TokenChainListResponse
		if err = json.NewDecoder(resp.Body).Decode(&tokenChains); err != nil {
			panic(errors.Wrap(err, "failed to unmarshal chain list response"))
		}

		if len(tokenChains.Data) == 0 {
			panic("at least 1 supported chain must be fetched from swapica-svc")
		}
		return tokenChains.Data
	}).([]resources.TokenChain)
}

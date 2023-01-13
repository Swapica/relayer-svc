package config

import (
	"crypto/ecdsa"

	"github.com/Swapica/relayer-svc/internal/signature"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (c *config) Transactor() signature.Signer {
	return c.transactor.Do(func() interface{} {
		var cfg struct {
			*ecdsa.PrivateKey `fig:"private_key,required"`
		}

		err := figure.
			Out(&cfg).
			With(figure.EthereumHooks).
			From(kv.MustGetStringMap(c.getter, "transactor")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out transactor"))
		}

		return signature.NewSigner(cfg.PrivateKey)
	}).(signature.Signer)
}

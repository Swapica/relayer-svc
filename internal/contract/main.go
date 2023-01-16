package contract

import (
	"math/big"

	"github.com/Swapica/relayer-svc/internal/config"
	"github.com/Swapica/relayer-svc/internal/signature"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func Transact(s signature.Signer, ch config.Chain) error {
	cli, err := ethclient.Dial(ch.RPC)
	if err != nil {
		return errors.Wrap(err, "failed to connect network by RPC")
	}

	t, err := NewSwapicaTransactor(ch.Contract, cli)
	if err != nil {
		return errors.Wrap(err, "failed to get contract transactor")
	}

	opts, _ := s.Opts(new(big.Int).SetInt64(ch.ID))
	_, err = t.ExecuteOrder(opts, []byte{}, [][]byte{})
	return errors.Wrap(err, "failed to call contract")
}

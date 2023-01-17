package tx

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/Swapica/relayer-svc/internal/gobind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type Transactorer interface {
	Transactor() Transactor
}

type Transactor interface {
	Transact(Chain) error
}

type transactor struct {
	privKey *ecdsa.PrivateKey
}

func NewTransactor(pk *ecdsa.PrivateKey) Transactor {
	return &transactor{
		privKey: pk,
	}
}

func (t *transactor) Transact(ch Chain) error {
	chainFields := map[string]interface{}{"chain_id": ch.ID, "chain_name": ch.Name}
	cli, err := ethclient.Dial(ch.RPC)
	if err != nil {
		return errors.Wrap(err, "failed to connect network by RPC", chainFields)
	}

	// FIXME this is the mock, replace it as soon as relayer contract is available
	tr, err := gobind.NewSwapicaTransactor(ch.Contract, cli)
	if err != nil {
		return errors.Wrap(err, "failed to get contract transactor", chainFields)
	}

	opts, _ := bind.NewKeyedTransactorWithChainID(t.privKey, new(big.Int).SetInt64(ch.ID))
	_, err = tr.ExecuteOrder(opts, []byte{}, [][]byte{})
	return errors.Wrap(err, "failed to call contract", chainFields)
}

package tx

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/crypto"
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
	Transact(Chain, []byte) error
}

type transactor struct {
	privKey *ecdsa.PrivateKey
}

func NewTransactor(pk *ecdsa.PrivateKey) Transactor {
	return &transactor{
		privKey: pk,
	}
}

func (t *transactor) Transact(ch Chain, data []byte) error {
	chainFields := map[string]interface{}{"chain_id": ch.ID, "chain_name": ch.Name}
	cli, err := ethclient.Dial(ch.RPC)
	if err != nil {
		return errors.Wrap(err, "failed to connect network by RPC", chainFields)
	}

	tr, err := gobind.NewRelayerTransactor(ch.Contract, cli)
	if err != nil {
		return errors.Wrap(err, "failed to get contract transactor", chainFields)
	}

	hash := crypto.Keccak256Hash(data)
	signature, err := crypto.Sign(hash.Bytes(), t.privKey)
	if err != nil {
		return errors.Wrap(err, "failed to sign transaction data", chainFields)
	}

	gasLimit, err := cli.EstimateGas(context.Background(), ethereum.CallMsg{})
	if err != nil {
		return errors.Wrap(err, "failed to estimate gas", chainFields)
	}

	opts, _ := bind.NewKeyedTransactorWithChainID(t.privKey, new(big.Int).SetInt64(ch.ID))
	opts.GasLimit = gasLimit

	_, err = tr.Execute(opts, data, [][]byte{signature})
	return errors.Wrap(err, "failed to call contract", chainFields)
}

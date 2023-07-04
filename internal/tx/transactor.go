package tx

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/Swapica/relayer-svc/internal/gobind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	sha3 "github.com/miguelmota/go-solidity-sha3"
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

	opts, _ := bind.NewKeyedTransactorWithChainID(t.privKey, new(big.Int).SetInt64(ch.ID))
	opts.GasLimit = 300000

	signature, err := crypto.Sign(sha3.SoliditySHA3(
		sha3.String("\x19Ethereum Signed Message:\n32"),
		sha3.Bytes32(data),
	), t.privKey)
	if err != nil {
		return errors.Wrap(err, "failed to sign transaction data", chainFields)
	}
	signature[64] += 27

	_, err = tr.Execute(opts, data, [][]byte{signature})
	return errors.Wrap(err, "failed to call contract", chainFields)
}

func signHash(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg))
}

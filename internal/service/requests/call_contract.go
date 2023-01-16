package requests

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Swapica/relayer-svc/resources"
	val "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type EvmTransactionRequest struct {
	Data    resources.EvmTransaction `json:"data"`
	ChainID int64
}

func NewCallContractRequest(r *http.Request) (EvmTransactionRequest, error) {
	var dst EvmTransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&dst); err != nil {
		return dst, errors.Wrap(err, "failed to decode request body")
	}

	cid, err := strconv.ParseInt(dst.Data.Attributes.ChainId, 16, 64)
	if err != nil {
		return dst, val.Errors{"data/attributes/chain_id": errors.Wrap(err, "failed to parse int")}
	}
	dst.ChainID = cid

	return dst, val.Errors{
		"data/type": val.Validate(dst.Data.Type, val.Required, val.In(resources.EVM_TRANSACTION)),
	}.Filter()
}

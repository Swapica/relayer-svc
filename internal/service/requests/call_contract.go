package requests

import (
	"encoding/json"
	"net/http"

	"github.com/Swapica/relayer-svc/resources"
	val "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func NewCallContractRequest(r *http.Request) (resources.EvmTransactionRequest, error) {
	var dst resources.EvmTransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&dst); err != nil {
		return dst, errors.Wrap(err, "failed to decode request body")
	}

	return dst, val.Errors{
		"data/type": val.Validate(dst.Data.Type, val.Required, val.In(resources.EVM_TRANSACTION)),
	}.Filter()
}

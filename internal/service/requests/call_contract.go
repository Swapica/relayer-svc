package requests

import (
	"encoding/json"
	"github.com/Swapica/relayer-svc/resources"
	val "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
	"regexp"
)

var hexRegexp = regexp.MustCompile("^0x[0-9a-f]+$")
var hexLengthRule = func(value interface{}) error {
	data := value.(string)
	if len(data)%2 != 0 {
		return errors.New("data length must be even")
	}
	return nil
}

func NewCallContractRequest(r *http.Request) (resources.EvmTransactionRequest, error) {
	var dst resources.EvmTransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&dst); err != nil {
		return dst, errors.Wrap(err, "failed to decode request body")
	}

	return dst, val.Errors{
		"data/type":             val.Validate(dst.Data.Type, val.Required, val.In(resources.EVM_TRANSACTION)),
		"data/attributes/data":  val.Validate(dst.Data.Attributes.Data, val.Required, val.Match(hexRegexp), val.By(hexLengthRule)),
		"data/attributes/chain": val.Validate(dst.Data.Attributes.Chain, val.Required),
	}.Filter()
}

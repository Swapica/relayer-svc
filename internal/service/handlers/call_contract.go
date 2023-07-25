package handlers

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"gitlab.com/distributed_lab/logan/v3"
	"net/http"

	"github.com/Swapica/relayer-svc/internal/service/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func CallContract(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCallContractRequest(r)
	if err != nil {
		Log(r).WithError(err).Debug("bad request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	chain := Chains(r).Get(request.Data.Attributes.Chain)
	if chain == nil {
		Log(r).WithField("chain", request.Data.Attributes.Chain).Debug("non-existent chain name")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	tChain := TokenChains(r).Get("0x"+request.Data.Attributes.Data[90:130], request.Data.Attributes.Chain)
	if tChain == nil {
		Log(r).WithFields(logan.F{
			"token": "0x" + request.Data.Attributes.Data[90:130],
			"chain": request.Data.Attributes.Chain,
		}).Debug("non-existent token chain")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	data, err := hexutil.Decode(request.Data.Attributes.Data)
	if err != nil {
		Log(r).WithError(err).Error("failed to decode transaction data")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if err = Transactor(r).Transact(*chain, data); err != nil {
		Log(r).WithError(err).Error("failed to send transaction")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

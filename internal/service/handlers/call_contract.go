package handlers

import (
	"net/http"

	"github.com/Swapica/relayer-svc/internal/contract"
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

	chain := Chains(r).Get(*request.Data.Attributes.ChainId)
	if chain == nil {
		Log(r).Debug("non-existent chain ID")
		ape.RenderErr(w, problems.BadRequest(nil)...)
		return
	}

	if err = contract.Transact(Transactor(r), *chain); err != nil {
		Log(r).WithError(err).Error("failed to send transaction")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

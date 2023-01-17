package handlers

import (
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

	if err = Transactor(r).Transact(*chain); err != nil {
		Log(r).WithError(err).Error("failed to send transaction")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

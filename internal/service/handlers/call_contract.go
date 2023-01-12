package handlers

import (
	"github.com/Swapica/relayer-svc/internal/service/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
)

func CallContract(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCallContractRequest(r)
	if err != nil {
		Log(r).WithError(err).Debug("bad request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	_ = request

	w.WriteHeader(http.StatusNoContent)
}

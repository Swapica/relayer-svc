package handlers

import (
	"context"
	"net/http"

	"github.com/Swapica/relayer-svc/internal/tx"
	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	chainsCtxKey
	transactorCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxChains(chains tx.Chains) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, chainsCtxKey, chains)
	}
}

func Chains(r *http.Request) tx.Chains {
	return r.Context().Value(chainsCtxKey).(tx.Chains)
}

func CtxTransactor(signer tx.Transactor) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, transactorCtxKey, signer)
	}
}

func Transactor(r *http.Request) tx.Transactor {
	return r.Context().Value(transactorCtxKey).(tx.Transactor)
}

package handlers

import (
	"context"
	"net/http"

	"github.com/Swapica/relayer-svc/internal/config"
	"github.com/Swapica/relayer-svc/internal/signature"
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

func CtxChains(chains config.Chains) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, chainsCtxKey, chains)
	}
}

func Chains(r *http.Request) config.Chains {
	return r.Context().Value(chainsCtxKey).(config.Chains)
}

func CtxTransactor(signer signature.Signer) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, transactorCtxKey, signer)
	}
}

func Transactor(r *http.Request) signature.Signer {
	return r.Context().Value(transactorCtxKey).(signature.Signer)
}

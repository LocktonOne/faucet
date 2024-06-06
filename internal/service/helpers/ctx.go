package helpers

import (
	"context"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/tokene/doorman/connector"
	"gitlab.com/tokene/faucet/internal/config"
	"gitlab.com/tokene/faucet/internal/signature"

	"net/http"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	authConfigCtxKey
	ethrpcConfigCtxKey
	signerCtxKey
	doormanConnectorCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxEthRPCConfig(entry *config.EthRPCConfig) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, ethrpcConfigCtxKey, entry)
	}
}

func EthRPCConfig(r *http.Request) *config.EthRPCConfig {
	return r.Context().Value(ethrpcConfigCtxKey).(*config.EthRPCConfig)
}

func CtxSigner(entry signature.Signer) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, signerCtxKey, entry)
	}
}

func Signer(r *http.Request) signature.Signer {
	return r.Context().Value(signerCtxKey).(signature.Signer)
}

func DoormanConnector(r *http.Request) connector.ConnectorI {
	return r.Context().Value(doormanConnectorCtxKey).(connector.ConnectorI)
}

func CtxDoormanConnector(entry connector.ConnectorI) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, doormanConnectorCtxKey, entry)
	}
}

func Token(r *http.Request) string {
	return r.Context().Value("token").(string)
}

func Address(r *http.Request) string {
	return r.Context().Value("address").(string)
}

package helpers

import (
	"context"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/tokene/doorman/connector"
	"gitlab.com/tokene/faucet/internal/config"

	"net/http"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	authConfigCtxKey
	ethrpcConfigCtxKey
	senderConfigCtxKey
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

func CtxAuthConfig(entry *config.AuthConfig) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, authConfigCtxKey, entry)
	}
}

func AuthConfig(r *http.Request) *config.AuthConfig {
	return r.Context().Value(authConfigCtxKey).(*config.AuthConfig)
}

func CtxEthRPCConfig(entry *config.EthRPCConfig) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, ethrpcConfigCtxKey, entry)
	}
}

func EthRPCConfig(r *http.Request) *config.EthRPCConfig {
	return r.Context().Value(ethrpcConfigCtxKey).(*config.EthRPCConfig)
}

func CtxSenderRPCConfig(entry *config.SenderConfig) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, senderConfigCtxKey, entry)
	}
}

func SenderRPCConfig(r *http.Request) *config.SenderConfig {
	return r.Context().Value(senderConfigCtxKey).(*config.SenderConfig)
}

func CtxDoormanConnector(entry connector.ConnectorI) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, doormanConnectorCtxKey, entry)
	}
}

func DoormanConnector(r *http.Request) connector.ConnectorI {
	return r.Context().Value(doormanConnectorCtxKey).(connector.ConnectorI)
}

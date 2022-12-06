package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/copus"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/kit/kv"
	doormanCfg "gitlab.com/tokene/doorman/connector/config"
	"gitlab.com/tokene/faucet/internal/signature"
)

type Config interface {
	comfig.Logger
	types.Copuser
	comfig.Listenerer
	AuthConfiger
	signature.Signerer
	EthRPCConfiger
	doormanCfg.DoormanConfiger
}

type config struct {
	comfig.Logger
	types.Copuser
	comfig.Listenerer
	AuthConfiger
	EthRPCConfiger
	signature.Signerer
	getter kv.Getter

	doormanCfg.DoormanConfiger
}

func New(getter kv.Getter) Config {
	return &config{
		getter:          getter,
		DoormanConfiger: doormanCfg.NewDoormanConfiger(getter),
		Signerer:        signature.NewSignerer(getter),
		EthRPCConfiger:  NewEthRPCConfiger(getter),
		AuthConfiger:    NewAuthConfiger(getter),
		Copuser:         copus.NewCopuser(getter),
		Listenerer:      comfig.NewListenerer(getter),
		Logger:          comfig.NewLogger(getter, comfig.LoggerOpts{}),
	}
}

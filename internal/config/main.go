package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/copus"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/kit/kv"
	doormanCfg "gitlab.com/tokene/doorman/connector/config"
)

type Config interface {
	comfig.Logger
	types.Copuser
	comfig.Listenerer
	SenderConfiger
	AuthConfiger
	doormanCfg.DoormanConfiger
	EthRPCConfiger
}

type config struct {
	comfig.Logger
	types.Copuser
	comfig.Listenerer
	AuthConfiger
	SenderConfiger
	EthRPCConfiger
	doormanCfg.DoormanConfiger
	getter kv.Getter
}

func New(getter kv.Getter) Config {
	return &config{
		getter:          getter,
		EthRPCConfiger:  NewEthRPCConfiger(getter),
		SenderConfiger:  NewSenderConfiger(getter),
		AuthConfiger:    NewAuthConfiger(getter),
		Copuser:         copus.NewCopuser(getter),
		Listenerer:      comfig.NewListenerer(getter),
		DoormanConfiger: doormanCfg.NewDoormanConfiger(getter),
		Logger:          comfig.NewLogger(getter, comfig.LoggerOpts{}),
	}
}

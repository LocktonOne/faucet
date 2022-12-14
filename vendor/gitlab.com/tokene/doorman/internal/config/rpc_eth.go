package config

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/url"
)

type EthRPCConfiger interface {
	EthRPCConfig() *EthRPCConfig
}

type EthRPCConfig struct {
	Endpoint string `fig:"endpoint"`
}

func NewEthRPCConfiger(getter kv.Getter) EthRPCConfiger {
	return &ethRPCConfig{
		getter: getter,
	}
}

type ethRPCConfig struct {
	getter kv.Getter
	once   comfig.Once
}

func (c *ethRPCConfig) EthRPCConfig() *EthRPCConfig {
	return c.once.Do(func() interface{} {
		raw := kv.MustGetStringMap(c.getter, "eth_rpc")
		config := EthRPCConfig{}
		err := figure.Out(&config).From(raw).Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out"))
		}

		return &config
	}).(*EthRPCConfig)
}

func (c *EthRPCConfig) EthRPCURL() *url.URL {
	u, err := url.Parse(c.Endpoint)
	if err != nil {
		panic(err)
	}
	return u
}

func (c *EthRPCConfig) EthClient() *ethclient.Client {
	client, err := ethclient.Dial(c.Endpoint)
	if err != nil {
		return nil
	}
	return client
}

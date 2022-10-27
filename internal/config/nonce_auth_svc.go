package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type AuthConfiger interface {
	AuthConfig() *AuthConfig
}

type AuthConfig struct {
	Endpoint string `fig:"endpoint"`
}

func NewAuthConfiger(getter kv.Getter) AuthConfiger {
	return &authConfig{
		getter: getter,
	}
}

type authConfig struct {
	getter kv.Getter
	once   comfig.Once
}

func (c *authConfig) AuthConfig() *AuthConfig {
	return c.once.Do(func() interface{} {
		raw := kv.MustGetStringMap(c.getter, "eth_ws")
		config := AuthConfig{}
		err := figure.Out(&config).From(raw).Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out"))
		}

		return &config
	}).(*AuthConfig)
}

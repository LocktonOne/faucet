package config

import (
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/url"
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
		raw := kv.MustGetStringMap(c.getter, "nonce_auth_svc")
		config := AuthConfig{}
		err := figure.Out(&config).From(raw).Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out"))
		}

		return &config
	}).(*AuthConfig)
}

func (c *authConfig) AuthURL() *url.URL {
	u, err := url.Parse(c.AuthConfig().Endpoint)
	if err != nil {
		panic(err)
	}
	return u
}

package config

import (
	credential "smartm2m-technical-test/internal/pkg/credential/config"
	ethclient "smartm2m-technical-test/internal/pkg/ethclient/config"
	httpserver "smartm2m-technical-test/internal/pkg/http/server/config"
)

type Config struct {
	Stage      string            `mapstructure:"stage"`
	Origins    []string          `mapstructure:"origins"`
	Domain     string            `mapstructure:"domain"`
	Creds      credential.Config `mapstructure:"creds"`
	HttpServer httpserver.Config `mapstructure:"httpserver"`
	EthClient  ethclient.Config  `mapstructure:"ethclient"`
}

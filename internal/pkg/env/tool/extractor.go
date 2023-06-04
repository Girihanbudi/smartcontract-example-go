package tool

import (
	credential "smartm2m-technical-test/internal/pkg/credential/config"
	"smartm2m-technical-test/internal/pkg/env/config"
	ethclient "smartm2m-technical-test/internal/pkg/ethclient/config"
	httpserver "smartm2m-technical-test/internal/pkg/http/server/config"
)

func ExtractCredsConfig(config config.Config) credential.Config {
	return config.Creds
}

func ExtractHttpServerConfig(config config.Config) httpserver.Config {
	return config.HttpServer
}

func ExtractEthClientConfig(config config.Config) ethclient.Config {
	return config.EthClient
}

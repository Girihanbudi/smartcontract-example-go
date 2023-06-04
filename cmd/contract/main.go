package main

import (
	"context"
	"fmt"
	"smartm2m-technical-test/internal/app/filesystem"
	"smartm2m-technical-test/internal/pkg/env"
	"smartm2m-technical-test/internal/pkg/ethclient"
	"smartm2m-technical-test/internal/pkg/log"
)

var Instance = "Smart Contract Deployment"

func main() {
	// Init app environment
	defaultEnvOps := env.NewDefaultOptions()
	env.InitEnv(defaultEnvOps)

	// Create app context
	ctx := context.Background()
	config := env.ProvideEnv().EthClient
	ethClient := ethclient.NewClient(ethclient.Options{Config: config})

	// Get trx opts for authentication
	trxOpts, err := ethClient.GetTrxOpts(ctx, config.SmHolder)
	if err != nil {
		log.Fatal(Instance, "failed to get trx opts", err)
	}

	ethClient.SetWithDefaultTrxOptsVal(trxOpts)

	// Deploying smart contract
	smAddress, _, _, err := filesystem.DeployFilesystem(trxOpts, ethClient.Client) //api is redirected from api directory from our contract go file
	if err != nil {
		log.Fatal(Instance, "failed to deploy file system contract", err)
	}

	// Log Smart Contract Address
	log.Print(Instance, fmt.Sprintf("Your Smart Contract Address is '%s', copy it to config 'smartcontractaddress' field", smAddress))
}

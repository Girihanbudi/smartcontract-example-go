package ethclient

import (
	"fmt"
	"smartm2m-technical-test/internal/pkg/ethclient/config"
	"smartm2m-technical-test/internal/pkg/log"

	goeth "github.com/ethereum/go-ethereum/ethclient"
)

const Instance = "Eth Client"

type Options struct {
	config.Config
}

type Client struct {
	Options
	Client *goeth.Client
}

func NewClient(options Options) *Client {

	client, err := goeth.Dial(options.Host)
	if err != nil {
		log.Fatal(Instance, "failed to init connection with provider", err)
	}

	log.Print(Instance, fmt.Sprintf("connected to %s", options.Host))

	return &Client{
		Options: options,
		Client:  client,
	}
}

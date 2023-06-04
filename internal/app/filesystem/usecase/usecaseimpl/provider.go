package usecaseimpl

import "smartm2m-technical-test/internal/pkg/ethclient"

type Options struct {
	EthClient *ethclient.Client
}

type Usecase struct {
	Options
}

func NewfilesystemUsecase(options Options) *Usecase {
	return &Usecase{options}
}

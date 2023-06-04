package usecaseimpl

import (
	"context"
	"fmt"
	module "smartm2m-technical-test/internal/app/filesystem"
	"smartm2m-technical-test/internal/pkg/stderror"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (u Usecase) GetTrxOpts(ctx context.Context, address string) (trxOpts *bind.TransactOpts, err *stderror.StdError) {
	// Get trx opts as credential
	trxOpts, getTrxOptsErr := u.EthClient.GetTrxOpts(ctx, address)
	if getTrxOptsErr != nil {
		err = stderror.ErrBadRequest().SetError(fmt.Errorf("address is not valid, %s", getTrxOptsErr.Error()))
		return
	}

	u.EthClient.SetWithDefaultTrxOptsVal(trxOpts)

	return
}

func (u Usecase) GetTrxOptsAndConn(ctx context.Context, address string) (trxOpts *bind.TransactOpts, conn *module.Filesystem, err *stderror.StdError) {
	// Get trx opts as credential
	trxOpts, err = u.GetTrxOpts(ctx, address)
	if err != nil {
		return
	}

	// Create connection
	conn, createConnErr := module.NewFilesystem(common.HexToAddress(u.EthClient.SmAddr), u.EthClient.Client)
	if createConnErr != nil {
		err = stderror.ErrBadRequest().SetError(fmt.Errorf("failed to create connection, %s", createConnErr.Error()))
		return
	}

	return
}

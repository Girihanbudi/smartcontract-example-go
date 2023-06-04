package usecaseimpl

import (
	"context"
	"errors"
	"smartm2m-technical-test/internal/app/filesystem/preset/request"
	"smartm2m-technical-test/internal/app/filesystem/preset/response"
	"smartm2m-technical-test/internal/pkg/stderror"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func (u Usecase) Download(ctx context.Context, cmd request.Download) (res response.Download, err *stderror.StdError) {
	// Validate input command
	if _, err = cmd.Validate(); err != nil {
		return
	}

	// Get connection
	trxOpts, conn, err := u.GetTrxOptsAndConn(ctx, cmd.Address)
	if err != nil {
		return
	}

	// Get file stored in smart contract
	smFile, getSmFileErr := conn.GetFile(&bind.CallOpts{From: trxOpts.From})
	if getSmFileErr != nil {
		err = stderror.ErrBadRequest().SetError(getSmFileErr)
		return
	}

	if smFile.Name == "" {
		err = stderror.ErrBadRequest().SetError(errors.New("file not found"))
	}

	res.Name = smFile.Name
	res.File = []byte(smFile.File)

	return
}

package usecase

import (
	"context"
	"smartm2m-technical-test/internal/app/filesystem/preset/request"
	"smartm2m-technical-test/internal/app/filesystem/preset/response"
	"smartm2m-technical-test/internal/pkg/stderror"
)

type Ifilesystem interface {
	Upload(ctx context.Context, cmd request.Upload) (res response.Upload, err *stderror.StdError)
	Download(ctx context.Context, cmd request.Download) (res response.Download, err *stderror.StdError)
}

package rest

import (
	"net/http"
	"smartm2m-technical-test/internal/app/filesystem/preset/request"
	"smartm2m-technical-test/internal/pkg/stdresponse"

	"github.com/gin-gonic/gin"
)

// Upload
//
//	@Summary			Upload file to smart contract.
//	@Tags					filesytem
//	@Router				/filesytem [post]
func (h Handler) Upload(ctx *gin.Context) {
	var req request.Upload
	ctx.ShouldBind(&req)

	if res, err := h.Filesystem.Upload(ctx, req); err != nil {
		stdresponse.GinMakeHttpResponseErr(ctx, err)
	} else {
		stdresponse.GinMakeHttpResponse(ctx, http.StatusCreated, res, nil)
	}
}

// Download
//
//	@Summary			Download file from smart contract.
//	@Tags					filesytem
//	@Router				/filesytem [get]
func (h Handler) Download(ctx *gin.Context) {
	var req request.Download
	ctx.ShouldBind(&req)
	if res, err := h.Filesystem.Download(ctx, req); err != nil {
		stdresponse.GinMakeHttpResponseErr(ctx, err)
	} else {

		stdresponse.GinMakeHttpResponseFile(ctx, http.StatusOK, res.Name, res.File)
	}
}

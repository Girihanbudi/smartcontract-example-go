package stdresponse

import (
	"fmt"
	"smartm2m-technical-test/internal/pkg/stderror"

	"github.com/gin-gonic/gin"
)

type ResponseBody struct {
	Data     interface{}        `json:"data,omitempty"`
	Metadata interface{}        `json:"metadata,omitempty"`
	Error    *stderror.StdError `json:"error,omitempty"`
}

func GinMakeHttpResponse(ctx *gin.Context, code int, result interface{}, metadata interface{}) {
	ctx.JSON(code, ResponseBody{
		Data:     result,
		Metadata: metadata,
	})
}

func GinMakeHttpResponseFile(ctx *gin.Context, code int, name string, file []byte) {
	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", name))
	ctx.Data(code, "application/octet-stream", file)
}

func GinMakeHttpResponseErr(ctx *gin.Context, err *stderror.StdError) {
	ctx.AbortWithStatusJSON(err.Status, ResponseBody{
		Error: err,
	})
}

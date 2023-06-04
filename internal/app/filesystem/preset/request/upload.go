package request

import (
	"fmt"
	"mime/multipart"
	module "smartm2m-technical-test/internal/app/filesystem"
	"smartm2m-technical-test/internal/pkg/io"
	"smartm2m-technical-test/internal/pkg/stderror"
)

type Upload struct {
	Address string                `form:"address"`
	File    *multipart.FileHeader `form:"file"`
}

func (req *Upload) Validate() (valid bool, err *stderror.StdError) {
	var errors []any

	if req.Address == "" {
		errors = append(errors, stderror.StdInvalidParams{
			Name:   "address",
			Reason: "Field is required",
		})
	}

	if req.File == nil {
		errors = append(errors, stderror.StdInvalidParams{
			Name:   "file",
			Reason: "Field is required",
		})
	} else if exist := io.ExtensionExist(req.File.Filename, module.AllowedExtensions...); !exist {
		errors = append(errors, stderror.StdInvalidParams{
			Name:   "file",
			Reason: fmt.Sprintf("Allowed type is: %v", module.AllowedExtensions),
		})
	}

	if len(errors) > 0 {
		err = stderror.ErrBadRequest().AddErrors(errors...)
	} else {
		valid = true
	}

	return
}

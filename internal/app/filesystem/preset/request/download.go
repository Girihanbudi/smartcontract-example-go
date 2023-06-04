package request

import "smartm2m-technical-test/internal/pkg/stderror"

type Download struct {
	Address string `form:"address"`
}

func (req *Download) Validate() (valid bool, err *stderror.StdError) {
	var errors []any

	if req.Address == "" {
		errors = append(errors, stderror.StdInvalidParams{
			Name:   "address",
			Reason: "Field is required",
		})
	}

	if len(errors) > 0 {
		err = stderror.ErrBadRequest().AddErrors(errors...)
	} else {
		valid = true
	}

	return
}

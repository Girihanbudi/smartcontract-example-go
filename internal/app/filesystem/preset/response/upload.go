package response

import (
	module "smartm2m-technical-test/internal/app/filesystem"
)

type Upload struct {
	Files []module.File `json:"files"`
}

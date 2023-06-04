package wire

import (
	"smartm2m-technical-test/internal/app/filesystem/api/rest"
	"smartm2m-technical-test/internal/app/filesystem/usecase"
	"smartm2m-technical-test/internal/app/filesystem/usecase/usecaseimpl"

	"github.com/google/wire"
)

var ModuleSet = wire.NewSet(
	usecaseSet,
	apiSet,
)

var usecaseSet = wire.NewSet(
	wire.Struct(new(usecaseimpl.Options), "*"),
	usecaseimpl.NewfilesystemUsecase,
	wire.Bind(new(usecase.Ifilesystem), new(*usecaseimpl.Usecase)),
)

var apiSet = wire.NewSet(

	wire.Struct(new(rest.Options), "*"),
	rest.NewfilesystemHandler,
)

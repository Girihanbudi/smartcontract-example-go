package wire

import (
	"smartm2m-technical-test/internal/pkg/http/server"
	"smartm2m-technical-test/internal/pkg/http/server/router"

	"github.com/google/wire"
)

var PackageSet = wire.NewSet(
	router.NewRouter,

	wire.Struct(new(server.Options), "*"),
	server.NewServer,
)

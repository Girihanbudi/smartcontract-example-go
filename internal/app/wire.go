//go:build wireinject
// +build wireinject

package app

import (
	filesystem "smartm2m-technical-test/internal/app/filesystem/wire"
	creds "smartm2m-technical-test/internal/pkg/credential/wire"
	env "smartm2m-technical-test/internal/pkg/env/wire"
	ethclient "smartm2m-technical-test/internal/pkg/ethclient/wire"
	httpserver "smartm2m-technical-test/internal/pkg/http/server/wire"

	"github.com/google/wire"
)

var AppSet = wire.NewSet(
	wire.Struct(new(Options), "*"),
	wire.Struct(new(App), "*"),
)

func NewApp() (*App, error) {
	panic(
		wire.Build(
			env.PackageSet,
			creds.PackageSet,
			httpserver.PackageSet,
			ethclient.PackageSet,

			AppSet,

			filesystem.ModuleSet,
		),
	)
}

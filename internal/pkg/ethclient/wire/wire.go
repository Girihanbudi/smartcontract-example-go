package wire

import (
	"smartm2m-technical-test/internal/pkg/ethclient"

	"github.com/google/wire"
)

var PackageSet = wire.NewSet(
	wire.Struct(new(ethclient.Options), "*"),
	ethclient.NewClient,
)

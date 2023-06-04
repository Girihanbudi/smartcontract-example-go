package wire

import (
	"smartm2m-technical-test/internal/pkg/env"
	"smartm2m-technical-test/internal/pkg/env/tool"

	"github.com/google/wire"
)

var PackageSet = wire.NewSet(
	env.ProvideEnv,
	tool.ExtractCredsConfig,
	tool.ExtractHttpServerConfig,
	tool.ExtractEthClientConfig,
)

package wire

import (
	"smartm2m-technical-test/internal/pkg/credential"

	"github.com/google/wire"
)

var PackageSet = wire.NewSet(
	wire.Struct(new(credential.Options), "*"),
	credential.NewTLSCredentials,
)

package credential

import (
	"crypto/tls"
	"smartm2m-technical-test/internal/pkg/credential/config"
	"smartm2m-technical-test/internal/pkg/env"
)

type Options struct {
	config.Config
}

type TlsCredentials struct {
	Options
	TlsCerts  *[]tls.Certificate
	TlsConfig *tls.Config
}

func NewTLSCredentials(options Options) (creds TlsCredentials) {
	creds.Options = options

	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair(options.PublicCert, options.PrivateKey)
	if err != nil {
		return creds
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
	}
	if env.CONFIG.Stage == string(env.StageLocal) {
		config.ClientAuth = tls.NoClientCert
	} else {
		config.ClientAuth = tls.RequireAndVerifyClientCert
	}

	tlsCerts := []tls.Certificate{serverCert}
	creds.TlsCerts = &tlsCerts
	creds.TlsConfig = config

	return creds
}

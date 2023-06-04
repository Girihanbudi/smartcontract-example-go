package http

import (
	"net/http"
	"smartm2m-technical-test/internal/pkg/env"
)

func DefaultSameSite() http.SameSite {
	if env.CONFIG.Stage != string(env.StageLocal) {
		return http.SameSiteStrictMode
	} else {
		return http.SameSiteNoneMode
	}
}

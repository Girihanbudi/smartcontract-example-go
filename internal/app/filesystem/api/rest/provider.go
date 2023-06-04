package rest

import (
	"github.com/gin-gonic/gin"

	filesystemusecase "smartm2m-technical-test/internal/app/filesystem/usecase"
)

type Options struct {
	Router *gin.Engine

	Filesystem filesystemusecase.Ifilesystem
}

type Handler struct {
	Options
}

func NewfilesystemHandler(options Options) *Handler {
	return &Handler{options}
}

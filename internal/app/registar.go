package app

import (
	_ "smartm2m-technical-test/docs"

	swaggerfiles "github.com/swaggo/files"
	ginswagger "github.com/swaggo/gin-swagger"
)

// registerRestHandler will register all endpoint route related to rest transport.
func (a App) registerRestHandler() {
	// Register rest handler
	a.filesystemRestHandler.RegisterApi()

	// Register swagger documentation
	a.HttpServer.Router.GET("/docs/*any", ginswagger.WrapHandler(swaggerfiles.Handler))
}

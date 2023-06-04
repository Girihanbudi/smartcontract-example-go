package app

import (
	"context"
	filesystemrest "smartm2m-technical-test/internal/app/filesystem/api/rest"
	loggermid "smartm2m-technical-test/internal/app/middleware/log"
	"smartm2m-technical-test/internal/pkg/credential"
	"smartm2m-technical-test/internal/pkg/http/server"
	httprouter "smartm2m-technical-test/internal/pkg/http/server/router"
	"smartm2m-technical-test/internal/pkg/log"
	"sync"

	"github.com/gin-gonic/gin"
)

var Instance = "App"

type Options struct {
	TlsCreds   credential.TlsCredentials
	HttpServer *server.Server

	filesystemRestHandler *filesystemrest.Handler
}

type App struct {
	Options
}

// Run will trigger the application runmodules and stopmodules gracefully.
func (a App) Run(ctx context.Context) {
	a.runModules(ctx)
	a.stopModules()
}

// runModules will init all required process for application.
func (a App) runModules(ctx context.Context) {
	log.Print(Instance, "starting service and connections...")

	// Recover from panic
	a.HttpServer.Router.Use(gin.Recovery())

	// GIN apply CORS setting
	a.HttpServer.Router.Use(httprouter.DefaultCORSSetting())

	// Log request and response to console
	a.HttpServer.Router.Use(loggermid.LogRequestToConsole())

	// Register all routes
	a.registerRestHandler()

	// Start to serve asynchronously
	go func() {
		a.HttpServer.Start()
	}()

	<-ctx.Done()
}

// stopModules is a function to stop application gracefully.
func (a App) stopModules() {
	log.Print(Instance, "stopping service and connections...")

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		a.HttpServer.Stop()
	}()

	wg.Wait()
	log.Print(Instance, "successfully stopped services and connections")
}

package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"smartm2m-technical-test/internal/app"
	"smartm2m-technical-test/internal/pkg/env"
	"syscall"
)

// @title           Smartm2m Technical API
// @version         1.0
// @description     This backend service content a definition for filesystem using smart contract

// @contact.name   API Support
// @contact.email  ghanbudi@gmail.com

// @host      localhost
func main() {
	// init app environment
	defaultEnvOps := env.NewDefaultOptions()
	env.InitEnv(defaultEnvOps)

	// create app context
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// create service app
	serviceApp, err := app.NewApp()
	if err != nil {
		log.Panic(err)
	}

	serviceApp.Run(ctx)
	stop()
}

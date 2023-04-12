package main

import (
	"context"
	"production-snippets/internal/app"
	"production-snippets/internal/config"
	"production-snippets/internal/logging"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2

func main() {
	cfg := config.GetConfig()
	logger := logging.GetLogger(cfg)

	ctx := logging.ContextWithLogger(context.Background(), logger)

	// a, err := app.NewApp(cfg, &logger)
	a, err := app.NewApp(cfg, ctx)
	if err != nil {
		logging.LoggerFromContext(ctx).Fatal().Err(err)
	}

	logging.LoggerFromContext(ctx).Info().Msg("Running application")
	a.Run()
}

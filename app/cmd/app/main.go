package main

import (
	"production-snippets/internal/app"
	"production-snippets/internal/config"
	"production-snippets/internal/logging"
)

// @title Production site snippets whith docker, zerologer, psql, GORM
// @version 1.0

// @contact.name Aleksandr Lesik
// @contact.url https://github.com/alekslesik
// @contact.email alekslesik@gmail.com

// @host http://bitrix.fvds.ru/
// @BasePath /v2

func main() {
	cfg := config.GetConfig()
	logger := logging.GetLogger(cfg)

	a, err := app.NewApp(cfg, &logger)
	if err != nil {
		logger.Fatal().Err(err)
	}

	logger.Info().Msg("Running application")
	a.Run()
}

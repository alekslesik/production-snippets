package main

import (
	"context"
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
	logger := logging.GetLogger()

	a, err := app.NewApp(cfg, &logger)
	if err != nil {
		logger.Fatal().Err(err)
	}

	err = a.Run(context.Background())
	if err != nil {
		logger.Fatal().Err(err)
	}
}

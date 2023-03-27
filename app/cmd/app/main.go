package main

import (
	"log"
	"production-snippets/internal/app"
	"production-snippets/internal/config"
	"production-snippets/internal/logger"
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
	log.Print("Config inititalizing.")
	cfg := config.GetConfig()

	log.Print("Logger inititalizing.")
	logger := logger.GetLogger(cfg)

	// app, err := app.NewApp(cfg, &logger)
	// if err != nil {
	// 	logger.Fatal()
	// }
}

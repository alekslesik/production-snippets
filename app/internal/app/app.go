package app

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"production-snippets/internal/config"

	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
	"github.com/rs/zerolog"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs" // docs is generated by Swag CLI, you have to import it.
	"github.com/swaggo/http-swagger/v2"
)

// Create app struct
type App struct {
	cfg    *config.Config
	logger *zerolog.Logger
}

// Return new instance of app
func NewApp(config *config.Config, logger *zerolog.Logger) (App, error) {
	logger.Print("Router initializing")
	router := chi.NewRouter()

	logger.Print("Swager docs initializing")
	// Swagger handler
	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://golang.fvds.ru:6666/swagger/doc.json"), //The url pointing to API definition
	))

	app := App{
		cfg:    config,
		logger: logger,
	}

	return app, nil
}

func startHTTP(a *App) {
	a.logger.Info().Msg("Start HTTP")

	var listener net.Listener

	if a.cfg.Listen.Type == "sock" {
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			a.logger.Fatal().Err(err)
		}
		socketPath := path.Join(appDir, a.cfg.Listen.SocketFile)
		a.logger.Info().Msgf("Socket path: %s", socketPath)

		a.logger.Info().Msg("Create and listen unix socket")
		listener, err = net.Listen("unix", socketPath)
		if err != nil {
			a.logger.Fatal().Err(err)
		}
	} else {
		a.logger.Info().Msgf("bind application to host: %s and port: %s", a.cfg.Listen.BindIP, a.cfg.Listen.Port)
		var err error
		listener, err = net.Listen("tcp", fmt.Sprintf("%s:%s", a.cfg.Listen.BindIP, a.cfg.Listen.Port))
		if err != nil {
			a.logger.Fatal().Err(err)
		}
	}

	// c := cors.New(cors.Options{
	// 	AllowedMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodPut, http.MethodOptions, http.MethodDelete},
	// 	AllowedOrigins:     []string{"http://localhost:6666", "http://localhost:8080"},
	// 	AllowCredentials:   true,
	// 	AllowedHeaders:     []string{"Location", "Charset", "Access-Control-Allow-Origin", "Content-Type", "content-type", "Origin", "Accept", "Content-Length", "Accept-Encoding", "X-CSRF-Token"},
	// 	OptionsPassthrough: true,
	// 	ExposedHeaders:     []string{"Location", "Authorization", "Content-Disposition"},
	// 	// Enable Debugging for testing, consider disabling in production
	// 	Debug: false,
	// })
	

}

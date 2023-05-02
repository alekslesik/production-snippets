package config

import (
	"flag"
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/ilyakaznacheev/cleanenv"
)

// Create config struct
type Config struct {
	IsDebug       bool `env:"IS_DEBUG" env-default:"false"`
	IsDevelopment bool `env:"IS_DEV" env-default:"true"`
	HTTP          struct {
		IP           string        `yaml:"ip" env:"HTTP_IP"`
		Port         int           `yaml:"ip" env:"HTTP_PORT"`
		ReadTimeOut  time.Duration `yaml:"ip" env:"HTTP_READ_TIMEOUT"`
		WriteTimeOut time.Duration `yaml:"ip" env:"HTTP_WRITE_TIMEOUT"`
		CORS         struct {
			AllowedMethods     []string `yaml:"allowed_methods" env:"HTTP_CORS_ALLOWED_METHODS"`
			AllowCredentials   bool     `yaml:"allow_credentials" env:"HTTP_CORS_ALLOW_CREDENTIALS"`
			AllowedOrigins     []string `yaml:"allowed_origins" env:"HTTP_CORS_ALLOWED_ORIGINS"`
			AllowedHeaders     []string `yaml:"allowed_headers"   env:"HTTP_CORS_ALLOWED_HEADERS"`
			OptionsPassthrough bool     `yaml:"options_passthrough" env:"HTTP_CORS_OPTIONS_PASSTHROUGH"`
			ExposedHeaders     []string `yaml:"exposed_headers" env:"HTTP_CORS_EXPOSED_HEADERS"`
			Debug              bool     `yaml:"debug" env:"HTTP_CORS_DEBUG"`
		} `yaml:"cors"`
	}
	AppConfig struct {
		LogLevel  string
		AdminUser struct {
			Email    string `env:"ADMIN_EMAIL" env-default:"admin"`
			Password string `env:"ADMIN_PWD" env-default:"admin"`
		}
	}
	LoggerSruct struct {
		Filename   string `env:"LOG_FILENAME" env-default:"/root/go/src/github.com/alekslesik/production-snippets/log.log"`
		MaxSize    int    `env:"LOG_MAXSIZE" env-default:"100"`
		MaxBackups int    `env:"LOG_MAXBACKUP" env-default:"3"`
		MaxAge     int    `env:"LOG_MAXAGE" env-default:"24"`
		Compress   bool   `env:"LOG_COMPRESS" env-default:"true"`
	}
	PostgreSQL struct {
		DSN string `env:"PROD_SNIP_DB_DSN" env-default:"postgres://production_snippets:486464@localhost:5432/production_snippets"`
	}
}

const (
	EnvConfigPathName  = "CONFIG_PATH"
	FlagConfigPathName = "config"
)

var configPath string
var instance *Config
var once sync.Once

// Return instance of config
func GetConfig() *Config {
	once.Do(func() {

		flag.StringVar(&configPath, FlagConfigPathName, "configs/config.local.yaml", "this is app config file")
		flag.Parse()

		if configPath == "" {
			configPath = os.Getenv(EnvConfigPathName)
		}

		if configPath == "" {
			log.Fatal().Msg("config path is required")
		}

		log.Info().Msg("Config init")

		instance = &Config{}

		if err := cleanenv.ReadEnv(instance); err != nil {
			helpText := "Production snippets"
			help, _ := cleanenv.GetDescription(instance, &helpText)
			
			log.Info().Msg(help)
			log.Fatal().Err(err)
		}
	})

	return instance
}

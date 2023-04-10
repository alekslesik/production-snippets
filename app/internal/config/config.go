package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

// Create config struct
type Config struct {
	Listen struct {
		Type       string `env:"LISTEN_TYPE" env-default:"port" env-description:"Port or Sock"`
		BindIP     string `env:"BIND_IP" env-default:"0.0.0.0"`
		Port       string `env:"PORT" env-default:"10000"`
		SocketFile string `env:"SOCKET_FILE" env-default:"app.sock"`
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
		Username string `env:"PSQL_USERNAME" env-default:"production_snippets"`
		Password string `env:"PSQL_PASSWORD" env-default:"486464"`
		Host     string `env:"PSQL_HOST" env-default:"localhost"`
		Port     string `env:"PSQL_PORT" env-default:"5432"`
		Database string `env:"PSQL_DATABASE" env-default:"production_snippets"`
	}

	IsDebug       bool `env:"IS_DEBUG" env-default:"false"`
	IsDevelopment bool `env:"IS_DEV" env-default:"true"`
}

var instance *Config
var once sync.Once

// Return instance of config
func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}

		if err := cleanenv.ReadEnv(instance); err != nil {
			helpText := "Monolith Note System"
			help, _ := cleanenv.GetDescription(instance, &helpText)
			log.Print(help)
			log.Fatal(err)
		}
	})

	return instance
}

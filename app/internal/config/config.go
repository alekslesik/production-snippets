package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

// Create config
type Config struct {
	Listen struct {
		Type   string `env:"LISTEN_TYPE" env-default:"port"`
		BindIP string `env:"BIND_IP" env-default:"0.0.0.0"`
		Port   string `env:"PORT" env-default:"10000"`
		SocketFile string
	}
	AppConfig struct {
		LogLevel  string
		AdminUser struct {
			Email    string `env:"ADMIN_EMAIL" env-required:"true"`
			Password string `env:"ADMIN_PWD" env-required:"true"`
		}
	}
	LoggerSruct struct {
		Filename   string `env:"LOG_FILENAME" env-required:"true"`
		MaxSize    int    `env:"LOG_MAXSIZE" env-required:"true"`
		MaxBackups int    `env:"LOG_MAXBACKUP" env-required:"true"`
		MaxAge     int    `env:"LOG_MAXAGE" env-required:"true"`
		Compress   bool   `env:"LOG_COMPRESS" env-required:"true"`
	}
	IsDebug       bool `env:"IS_DEBUG" env-default:"false"`
	IsDevelopment bool `env:"IS_DEV" env-default:"true"`
}

var instance *Config
var once sync.Once

// Return instance of config
func GetConfig() *Config {
	once.Do(func() {
		log.Print("gather config")

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

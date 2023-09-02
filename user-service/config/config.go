package config

import (
	"github.com/joho/godotenv"
)

type AppConfig struct {
	App struct {
		Env string
	}
	Grpc struct {
		Port string
		Host string
	}
	Postgres struct {
		Host string
		Port string
		User string
		Pass string
		Name string
	}
}

var appConfig *AppConfig

func NewAppConfig() *AppConfig {
	// add config file path in .env
	_ = godotenv.Load()

	if appConfig == nil {
		appConfig = &AppConfig{}

		initPostgres(appConfig)
		initApp(appConfig)
		initGrpc(appConfig)
	}

	return appConfig
}

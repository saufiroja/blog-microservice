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
	UserService struct {
		Host string
		Port string
	}
}

var appConfig *AppConfig

func NewAppConfig() *AppConfig {
	// add config file path in .env
	_ = godotenv.Load()

	if appConfig == nil {
		appConfig = &AppConfig{}

		initApp(appConfig)
		initGrpc(appConfig)
		initUserServices(appConfig)
	}

	return appConfig
}

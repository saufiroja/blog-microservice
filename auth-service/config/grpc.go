package config

import (
	"os"
)

func initGrpc(conf *AppConfig) {
	port := os.Getenv("AUTH_PORT")
	host := os.Getenv("AUTH_HOST")

	conf.Grpc.Port = port
	conf.Grpc.Host = host
}

func initUserServices(conf *AppConfig) {
	userHost := os.Getenv("USER_HOST")
	userPort := os.Getenv("USER_PORT")

	conf.UserService.Host = userHost
	conf.UserService.Port = userPort
}

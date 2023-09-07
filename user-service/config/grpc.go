package config

import (
	"os"
)

func initGrpc(conf *AppConfig) {
	port := os.Getenv("USER_PORT")
	host := os.Getenv("USER_HOST")

	conf.Grpc.Port = port
	conf.Grpc.Host = host
}

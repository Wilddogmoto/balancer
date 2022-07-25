package main

import (
	"github.com/Wilddogmoto/balacer/config"
	"github.com/Wilddogmoto/balacer/logger"
	"github.com/Wilddogmoto/balacer/server"
)

func main() {

	// initialize logger
	logger.InitLogger()

	// initialize config project
	config.InitConfig()

	// start grpc server
	server.GrpcServer()

}

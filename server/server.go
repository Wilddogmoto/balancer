package balancer

import (
	"github.com/Wilddogmoto/balacer/logger"
	log "github.com/sirupsen/logrus"
)

func ServerGrpc() {

	var logg = logger.Log.WithFields(log.Fields{
		"module": "grpc_server",
	})

	logg.Warn("!!!!!!")

}

package server

import (
	"github.com/Wilddogmoto/balacer/config"
	"github.com/Wilddogmoto/balacer/logger"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"sync"
)

type GRPCServer struct { // implements definition from Proto
	UnimplementedBalancerServer
	addr string
}

func GrpcServer() {

	var (
		logg = logger.Log.WithFields(log.Fields{
			"module": "grpc_server",
		})
		wg sync.WaitGroup
	)

	for _, addr := range config.Params.GRPCPorts {

		wg.Add(1)

		go func(add string) {

			var (
				err      error
				listener net.Listener
			)
			defer wg.Done()

			logg.Infof("start grpc server on port: %s", add)

			gs := grpc.NewServer()
			reflection.Register(gs)
			RegisterBalancerServer(gs, &GRPCServer{addr: add})

			listener, err = net.Listen("tcp", add)
			if err != nil {
				logg.Fatalf("error on create listener for grpc server: %v", err)
			}
			if err = gs.Serve(listener); err != nil {
				logg.Fatalf("error on create grpc server: %v", err)
			}
		}(addr)
	}

	wg.Wait()
}

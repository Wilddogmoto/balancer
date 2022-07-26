package main

import (
	"context"
	"fmt"
	"github.com/Wilddogmoto/balacer/config"
	"github.com/Wilddogmoto/balacer/logger"
	"github.com/Wilddogmoto/balacer/server"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"time"
)

const (
	scheme      = "test"
	serviceName = "balancer"
	testUrl     = "http://s1.origin-cluster/video/123/xcg2djHckad.m3u8"
)

func main() {

	var (
		conn *grpc.ClientConn
		err  error
		resp *server.ResponseBody
	)

	logger.InitLogger()
	config.InitConfig()

	logg := logger.Log.WithFields(log.Fields{
		"module": "grpc_client",
	})

	conn, err = grpc.Dial(fmt.Sprintf("%s:///%s", scheme, serviceName),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`),
	)
	defer conn.Close()

	if err != nil {
		logg.Fatalf("error on grpc dial: %v", err)
	}

	client := server.NewBalancerClient(conn)

	var id int

	for {
		id++

		ctx, _ := context.WithTimeout(context.Background(), 2500*time.Millisecond)

		resp, err = client.GetUrl(ctx, &server.RequestBody{Increment: int64(id), Video: testUrl})
		if err != nil {
			logg.Fatalf("error request: %v", err)
		}

		logg.Infof("count call: %d URL: %s", id, resp.GetUrl())

		//time.Sleep(1 * time.Second)
	}
}

type (
	exampleResolverBuilder struct{}

	exampleResolver struct {
		target     resolver.Target
		cc         resolver.ClientConn
		addrsStore map[string][]string
	}
)

func init() {
	resolver.Register(&exampleResolverBuilder{})
}

func (*exampleResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &exampleResolver{
		target: target,
		cc:     cc,
		addrsStore: map[string][]string{
			serviceName: config.Params.GRPCPorts,
		},
	}
	r.start()
	return r, nil
}
func (*exampleResolverBuilder) Scheme() string { return scheme }

func (r *exampleResolver) start() {
	addrStrs := r.addrsStore[r.target.Endpoint]
	addrs := make([]resolver.Address, len(addrStrs))
	for i, s := range addrStrs {
		addrs[i] = resolver.Address{Addr: s}
	}
	r.cc.UpdateState(resolver.State{Addresses: addrs})
}
func (*exampleResolver) ResolveNow(o resolver.ResolveNowOptions) {}
func (*exampleResolver) Close()                                  {}

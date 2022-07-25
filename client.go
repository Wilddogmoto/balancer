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

//func main() {
//
//	var (
//		conn *grpc.ClientConn
//		err  error
//		resp *server.ResponseBody
//		//address = "static:///localhost:8080,localhost:8081,localhost:8082"
//	)
//
//	resolver.SetDefaultScheme("dns")
//
//	conn, err = grpc.Dial("dns://dns_server/localhost:44300",
//		grpc.WithTransportCredentials(insecure.NewCredentials()),
//		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
//	)
//	if err != nil {
//		log.Fatalf("error on grpc dial: %v", err)
//	}
//
//	client := server.NewBalancerClient(conn)
//
//	var count int
//
//	for {
//		count++
//
//		ctx, _ := context.WithTimeout(context.Background(), 2500*time.Millisecond)
//
//		resp, err = client.GetUrl(ctx, &server.RequestBody{Video: "http://s1.origin-cluster/video/123/xcg2djHckad.m3u8"})
//		if err != nil {
//			log.Fatalf("error request: %v", err)
//		}
//
//		log.Printf("count: %d URL: %s", count, resp.GetUrl())
//
//	}
//}

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

	resolver.SetDefaultScheme("dns")

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

		logg.Infof("count: %d URL: %s", id, resp.GetUrl())

		//time.Sleep(1 * time.Second)
	}
}

// Following is an example name resolver implementation. Read the name
// resolution example to learn more about it.

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

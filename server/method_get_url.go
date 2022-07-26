package server

import (
	"context"
	"github.com/Wilddogmoto/balacer/config"
	"github.com/Wilddogmoto/balacer/logger"
	log "github.com/sirupsen/logrus"
	"strings"
)

var count int

func (c *GRPCServer) GetUrl(ctx context.Context, in *RequestBody) (*ResponseBody, error) {

	var (
		logg = logger.Log.WithFields(log.Fields{
			"module": "method_get",
		})
	)

	logg.Warnf("request for port: %s id call: %d ", c.addr, in.GetIncrement())

	if in.GetVideo() == "" {
		return nil, nil
	}

	count++

	if count == 1 {
		return &ResponseBody{Url: in.GetVideo()}, nil
	}

	if count == 10 {
		count = 0
	}

	return &ResponseBody{Url: urlCreator(in.GetVideo())}, nil
}

func urlCreator(body string) (uri string) {

	var (
		http   = "http://" + config.Params.CDNHost
		server string
		path   string
	)

	blocks := strings.Split(body, "/")

	for e, block := range blocks {

		if e == 0 || e == 1 {
			continue
		}

		if e == 2 {
			server = strings.Split(block, ".")[0]
			path += "/" + server
			continue
		}

		path += "/" + block

	}

	uri = http + path

	return
}

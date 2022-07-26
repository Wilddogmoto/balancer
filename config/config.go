package config

import (
	"github.com/Wilddogmoto/balacer/logger"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

type Config struct {
	CDNHost   string   `json:"cdn_host"`
	GRPCPorts []string `json:"grpc_ports"`
}

var Params Config

func InitConfig() {

	var logg = logger.Log.WithFields(log.Fields{
		"module": "config",
	})

	if err := godotenv.Load(); err != nil {
		logg.Fatalf("error on loading .env file: %v", err)
	}

	ports := strings.Split(os.Getenv("GRPC_PORT"), ".")

	addres := make([]string, 0, len(ports))
	host := os.Getenv("HOST")

	for _, port := range ports {
		addres = append(addres, host+port)
	}

	Params = Config{
		CDNHost:   os.Getenv("CDN_HOST"),
		GRPCPorts: addres,
	}

	logg.Infof("initilize config: %+v", Params)

}

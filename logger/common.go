package logger

import (
	log "github.com/sirupsen/logrus"
	"os"
)

var Log *log.Entry

func InitLogger() {

	logger := log.New()
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2022-01-01 18:30:51",
		ForceColors:     true,
	})

	logger.SetLevel(log.DebugLevel)

	Log = logger.WithFields(log.Fields{})

	Log.Infof("logger init success")
}

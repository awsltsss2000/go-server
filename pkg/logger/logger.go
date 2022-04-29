package logger

import (
	"go-server/pkg/time"
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: time.Layout,
		// PrettyPrint:     true,
		DataKey: "extra",
	})

	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

package main

import (
	"remoteworkout/internal/infra/config"

	log "github.com/sirupsen/logrus"
)

const port = 8080

func initLogging() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
}

func main() {
	initLogging()
	log.Infof("Starting remote-workout back-end at port %d", port)
	ctx := config.ProvideAppContext()
	ctx.Server.Listen(port)
}

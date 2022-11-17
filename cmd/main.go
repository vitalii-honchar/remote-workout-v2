package main

import (
	log "github.com/sirupsen/logrus"
)

func initLogging() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
}

func main() {
	initLogging()
	log.Info("Starting remote-workout back-end")
}

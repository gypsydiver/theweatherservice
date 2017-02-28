package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/agave/go-agave/config"
	"github.com/gypsydiver/theweatherservice/services/geolocation/geoDB"
	"github.com/gypsydiver/theweatherservice/services/geolocation/util"
)

func init() {
	configureLogger()
	loadConfiguration()
	initGeoDB()
}

func configureLogger() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.ErrorLevel)
}

func loadConfiguration() {
	err := config.GetConfig(&util.Config, "config.yml")
	if err != nil {
		log.WithFields(log.Fields{
			"error":    err,
			"defaults": util.Config,
		}).Warn("Config error, using default configuration")
	}
}

func initGeoDB() {
	err := geoDB.Init()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Panic("Error loading Geolite2 database")
		return
	}
	log.Info("Geolite2 database successfully loaded")
}

func main() {
	//server
}

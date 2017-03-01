package main

import (
	"os"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/agave/go-agave/config"
	"github.com/gypsydiver/theweatherservice/services/geolocation/geoDB"
	"github.com/gypsydiver/theweatherservice/services/geolocation/util"
)

func init() {
	configureLogger()
	loadConfiguration()
	initGeoDB()
	go updateGeoDB()
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
	if err := geoDB.UpdateDB(); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Panic("Error loading Geolite2 database")
		return
	}
	log.Info("Geolite2 database successfully loaded")
}

func updateGeoDB() {
	for {
		oneMonth := time.Hour * 24 * 31
		waitTime := oneMonth * time.Duration(util.Config.IntervalUpdateDBInMonths)
		time.Sleep(waitTime)
		log.Info("Updating Geolite2 DB")
		if err := geoDB.UpdateDB(); err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Warn("Geolite2 DB update failed")
		} else {
			log.Info("Geolite2 DB has been updated")
		}
	}
}

func main() {
	//server
}

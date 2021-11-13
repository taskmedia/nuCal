package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/taskmedia/nuCal/pkg/http/rest"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("Starting nuCal")
}

func main() {
	router := rest.SetupRouter()
	router.Run("0.0.0.0:8080")
}

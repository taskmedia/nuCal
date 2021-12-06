package main

import (
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	"github.com/taskmedia/nuCal/pkg/http/rest"
	"github.com/taskmedia/nuCal/pkg/persistence"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("Starting nuCal")
}

// struct defines which environment variables are available
type NuCalEnvs struct {
	Path string `default:"./"`
}

func main() {
	// get environment variables
	var env NuCalEnvs
	err := envconfig.Process("nucal", &env)
	if err != nil {
		log.WithField("err", err.Error()).Fatal("could not load environment variables")
	}

	log.WithFields(log.Fields{
		"NU_PATH": env.Path,
	}).Info("environment variables")

	// forward environment variables
	persistence.PathTree = env.Path

	router := rest.SetupRouter()
	router.Run("0.0.0.0:8080")
}

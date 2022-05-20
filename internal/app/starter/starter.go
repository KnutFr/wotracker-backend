package starter

import (
	log "github.com/sirupsen/logrus"
	config "wotracker-back/configs"
	"wotracker-back/internal/app/api"
)

func RunApi() api.Server {
	log.Debug("loading config from configs folder")
	loadedConfig := config.ReadConfigFile("configs")
	log.Debugf("starting server on port %s", loadedConfig.PortServer)
	server := api.Server{}
	server.StartServer(loadedConfig)
	return server
}

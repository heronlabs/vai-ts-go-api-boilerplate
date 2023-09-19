package config

import (
	config "hello-world/src/application/configuration/models"
	"os"
)

func BootstrapConfig() config.ConfigModel {
	var configModel config.ConfigModel

	apiPort := os.Getenv("API_PORT")
	if apiPort == "" {
		configModel.ApiPort = 3000
	}

	return configModel
}

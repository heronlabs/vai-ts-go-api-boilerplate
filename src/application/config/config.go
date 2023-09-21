package config

import (
	"fmt"
	config "hello-world/src/application/config/models"
	"os"
)

func Config() config.ConfigModel {
	var configModel config.ConfigModel

	apiPort := os.Getenv("API_PORT")
	if apiPort == "" {
		configModel.ApiPort = "1337"
	} else {
		configModel.ApiPort = apiPort
	}

	fmt.Println("CONFIG OK!")

	return configModel
}

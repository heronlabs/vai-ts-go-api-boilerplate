package main

import (
	"hello-world/src/application/api"
	"hello-world/src/application/config"
	"hello-world/src/application/server"

	"github.com/gin-gonic/gin"
)

func main() {
	config := config.Config()
	ginApp := gin.Default()
	api.Api(ginApp)
	server.Server(config.ApiPort, ginApp)
}

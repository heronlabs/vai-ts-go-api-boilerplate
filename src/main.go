package main

import (
	"go-api-boilerplate/src/application/api"
	"go-api-boilerplate/src/application/config"
	"go-api-boilerplate/src/application/server"

	"github.com/gin-gonic/gin"
)

func main() {
	config := config.Config()
	ginApp := gin.Default()
	api.Api(ginApp)
	server.Server(config.ApiPort, ginApp)
}

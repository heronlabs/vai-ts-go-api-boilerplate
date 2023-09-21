package api

import (
	"fmt"
	healthCheckController "go-api-boilerplate/src/application/api/controllers/health-check"

	"github.com/gin-gonic/gin"
)

func Api(ginApp *gin.Engine) {
	ginApp.GET("/", healthCheckController.GetIndex)
	ginApp.POST("/", healthCheckController.PostWebHook)
	fmt.Println("API OK!")
}

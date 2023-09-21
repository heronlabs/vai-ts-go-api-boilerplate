package healthCheckController

import (
	"go-api-boilerplate/src/application/api/presenters"

	"github.com/gin-gonic/gin"
)

func GetIndex(ctx *gin.Context) {
	response := presenters.Envelope("Server running!")

	ctx.JSON(200, response)
}

func PostWebHook(ctx *gin.Context) {
	var req WebHookModel
	req.Method = "POST"
	ctx.BindJSON(&req.Content)

	response := presenters.Envelope(req)

	ctx.JSON(200, response)
}

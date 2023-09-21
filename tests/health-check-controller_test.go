package healthCheckController

import (
	"bytes"
	"encoding/json"
	healthCheckController "hello-world/src/application/api/controllers/health-check"
	"hello-world/src/application/api/presenters"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetIndex(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.GET("/", healthCheckController.GetIndex)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)
	var response presenters.JsonModel
	json.NewDecoder(w.Body).Decode(&response)

	var expectedResponsePayload = presenters.JsonModel{
		Payload: "Server running!",
	}
	assert.Equal(t, expectedResponsePayload, response)
}

func TestPostWebhook(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.POST("/", healthCheckController.PostWebHook)

	body := map[string]interface{}{
		"testOne": "hello world!",
	}

	jsonBytes, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/", bytes.NewReader(jsonBytes))
	router.ServeHTTP(w, req)
	var response presenters.JsonModel
	json.NewDecoder(w.Body).Decode(&response)

	var expectedResponsePayload = presenters.JsonModel{
		Payload: healthCheckController.WebHookModel{
			Method:  "POST",
			Content: body,
		},
	}

	assert.ObjectsAreEqual(expectedResponsePayload, response)
}

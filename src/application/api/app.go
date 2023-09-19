package app

import (
	"fmt"
	healthCheckController "hello-world/src/application/api/controllers/health-check"
	"net/http"
)

func BootstrapApp() {
	http.HandleFunc("/", healthCheckController.GetIndex)
	http.HandleFunc("/webhook", healthCheckController.PostWebHook)

	fmt.Println("Bootstrap app")
}

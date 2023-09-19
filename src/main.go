package main

import (
	"errors"
	"fmt"
	"hello-world/src/application/api"
	config "hello-world/src/application/configuration"
	"log"
	"net/http"
)

func main() {
	config := config.BootstrapConfig()

	api.BootstrapApp()

	fmt.Println("Server listen on port:", config.ApiPort)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.ApiPort), nil); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("HTTP server error: %v", err)
	}
}

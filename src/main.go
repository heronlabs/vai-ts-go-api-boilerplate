package main

import (
	"fmt"
	app "hello-world/src/application/api"
	config "hello-world/src/application/configuration"
	"net/http"
)

func main() {
	config := config.BootstrapConfig()

	app.BootstrapApp()

	fmt.Println("Server listen on port:", config.ApiPort)

	http.ListenAndServe(fmt.Sprintf(":%d", config.ApiPort), nil)
}

package healthCheckController

import (
	healthCheckRes "hello-world/src/application/api/controllers/health-check/res"
	"hello-world/src/application/api/presenters"
	"net/http"
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	message := presenters.JsonPresenter{
		Payload: "Server running!",
	}

	presenters.Envelope(message, w, r)
}

func PostWebHook(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	content := presenters.DecodeBody(w, r)
	if content == nil {
		return
	}

	response := presenters.JsonPresenter{
		Payload: healthCheckRes.WebHookModel{
			Method:  r.Method,
			Content: content,
		},
	}

	presenters.Envelope(response, w, r)
}

package presenters

import (
	"encoding/json"
	"net/http"
)

type JsonPresenter struct {
	Payload any `json:"payload"`
}

func Envelope(payload JsonPresenter, w http.ResponseWriter, r *http.Request) {

	jsonData, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")

	_, err := w.Write(jsonData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

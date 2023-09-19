package presenters

import (
	"encoding/json"
	"net/http"
)

type JsonPresenter struct {
	Payload any `json:"payload"`
}

func Envelope(message JsonPresenter, w http.ResponseWriter, r *http.Request) {

	jsonData, err := json.Marshal(message)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

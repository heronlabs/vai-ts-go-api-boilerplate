package presenters

import (
	"encoding/json"
	"net/http"
)

type JsonPresenter struct {
	Payload any `json:"payload"`
}

func Envelope(payload JsonPresenter, w http.ResponseWriter, r *http.Request) []byte {
	jsonData, _ := json.Marshal(payload)
	return jsonData
}

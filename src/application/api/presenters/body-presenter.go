package presenters

import (
	"encoding/json"
	"net/http"
)

func DecodeBody(w http.ResponseWriter, r *http.Request) any {
	var body any

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}

	return body
}

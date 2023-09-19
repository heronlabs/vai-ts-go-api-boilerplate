package presenters

import (
	"encoding/json"
	"net/http"
)

func DecodeBody(w http.ResponseWriter, r *http.Request) any {
	var p any

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}

	return p
}

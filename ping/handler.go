package ping

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type pingResponse struct {
	Message string `json:"message"`
}

// Handler handles GET requests to /ping
func Handler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	response := pingResponse{"pong"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&response)
}

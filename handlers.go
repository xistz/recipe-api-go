package main

import (
	"fmt"
	"net/http"
)

type pingResponse struct {
	Message string `json:"message"`
}

// PingHandler handles GET requests to /ping
func PingHandler(s Store) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, fmt.Sprintf("HTTP Method %s Not Allowed", r.Method), http.StatusMethodNotAllowed)
			return
		}
		res := pingResponse{}
		w.Header().Set("Content-Type", "application/json")

		err := s.Ping()
		if err != nil {
			res.Message = err.Error()

			respondJSON(w, http.StatusServiceUnavailable, &res)
			return
		}

		res.Message = "pong"
		respondJSON(w, http.StatusOK, &res)
	}
}

package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type pingResponse struct {
	Message string `json:"message"`
}

// PingHandler handles GET requests to /ping
func PingHandler(s Store) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

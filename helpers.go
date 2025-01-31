package main

import (
	"encoding/json"
	"net/http"
	"os"
)

// loads ENV or sets default values for the api
func getEnv() (string, string, string, string, string) {
	dbUser, ok := os.LookupEnv("DB_USER")
	if !ok {
		dbUser = "retailai"
	}
	dbPassword, ok := os.LookupEnv("DB_PASSWORD")
	if !ok {
		dbPassword = "zH4tAwEfMUL7x3nM"
	}
	dbAddress, ok := os.LookupEnv("DB_ADDRESS")
	if !ok {
		dbAddress = "db:3306"
	}
	dbName, ok := os.LookupEnv("DB_NAME")
	if !ok {
		dbName = "recipes"
	}
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	return port, dbUser, dbPassword, dbAddress, dbName
}

func respondJSON(w http.ResponseWriter, code int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(body)
}

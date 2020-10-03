package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	// get env variables
	port, dbUser, dbPassword, dbHost, dbPort, dbName := getEnv()

	addr := ":" + port

	dbPool, err := initMySQLDB(dbUser, dbPassword, dbHost, dbPort, dbName)
	if err != nil {
		log.Fatal(err)
	}

	store := NewMySQLStore(dbPool)

	r := httprouter.New()

	r.GET("/ping", PingHandler(store))

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	log.Println("Listening on port", port)
	log.Fatal(
		http.ListenAndServe(
			addr,
			handlers.RecoveryHandler()(loggedRouter),
		),
	)
}

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
	port, dbUser, dbPassword, dbAddress, dbName := getEnv()

	addr := ":" + port

	dbPool, err := initMySQLDB(dbUser, dbPassword, dbAddress, dbName)
	if err != nil {
		log.Fatal(err)
	}

	store := NewMySQLStore(dbPool)

	r := httprouter.New()

	r.GET("/ping", PingHandler(store))
	r.GET("/recipes", ListHandler(store))
	r.POST("/recipes", CreateHandler(store))
	r.GET("/recipes/:id", FindHandler(store))
	r.DELETE("/recipes/:id", DeleteHandler(store))
	r.PATCH("/recipes/:id", UpdateHandler(store))

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	log.Println("Listening on port", port)
	log.Fatal(
		http.ListenAndServe(
			addr,
			handlers.RecoveryHandler()(loggedRouter),
		),
	)
}

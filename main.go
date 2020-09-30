package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
)

func main() {
	// get env variables
	dsn, ok := os.LookupEnv("DSN")
	if !ok {
		dsn = "retailai:zH4tAwEfMUL7x3nM@(db:3306)/recipes"
	}
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}
	addr := ":" + port

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	r := httprouter.New()

	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(addr, r))
}

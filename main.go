package main

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	// get env variables
	dsn, ok := os.LookupEnv("DSN")
	if !ok {
		dsn = "retailai:zH4tAwEfMUL7x3nM@(db:3306)/recipes"
	}

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
}

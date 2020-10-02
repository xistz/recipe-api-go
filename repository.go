package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func initMySQLDB(dbUser, dbPassword, dbHost, dbPort, dbName string) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser, dbPassword, dbHost, dbPort, dbName,
	)

	dbPool, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	dbPool.SetMaxIdleConns(5)
	dbPool.SetMaxOpenConns(5)
	dbPool.SetConnMaxLifetime(1800)

	return dbPool, nil
}

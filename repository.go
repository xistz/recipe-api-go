package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func initMySQLDB(dbUser, dbPassword, dbAddress, dbName string) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?parseTime=true",
		dbUser, dbPassword, dbAddress, dbName,
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

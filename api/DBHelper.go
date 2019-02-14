package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type DatabaseConf struct {
	DBDriver string
	DBUser   string
	DBPass   string
	DBName   string
}

// DBConn retorna conexa com o dabatase
func DBConn(conf *DatabaseConf) (db *sql.DB) {
	dbDriver := conf.DBDriver
	dbUser := conf.DBUser
	dbPass := conf.DBPass
	dbName := conf.DBName
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

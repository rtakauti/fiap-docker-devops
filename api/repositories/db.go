package repositories

import (
	"database/sql"

	"../shared"
	_ "github.com/go-sql-driver/mysql"
)

// DBConn retorna conexa com o dabatase
func DBConn(conf *shared.DatabaseConf) (db *sql.DB) {
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

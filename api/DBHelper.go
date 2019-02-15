package main

import (
	"database/sql"
	"fmt"
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
	connectionUrl := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", conf.DBUser, conf.DBPass, "tcp", "mysql", 3306, conf.DBName)

	fmt.Println("Conexão com banco de dados: " + connectionUrl)

	db, err := sql.Open(conf.DBDriver, connectionUrl)

	if err != nil {
		panic(err.Error())
	}

	return db
}

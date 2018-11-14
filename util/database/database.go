package database

import (
	"bigproject/util/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func ConnectDB() (db *sql.DB) {
	conf := config.GetConfig()
	stringConfig := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", conf.DBUser, conf.DBPass, conf.DBHost, conf.DBPort, conf.DBName)
	dbConn, err := sql.Open("postgres", stringConfig)
	if err != nil {
		panic(err.Error())
	}
	return dbConn
}
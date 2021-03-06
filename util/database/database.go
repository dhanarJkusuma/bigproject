package database

import (
	"bigproject/util/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

func ConnectDB() (db *sql.DB) {
	conf := config.GetConfig()
	log.Printf("[BigProject] : Connecting to db %v \n", conf.DBName)
	stringConfig := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", conf.DBUser, conf.DBPass, conf.DBHost, conf.DBPort, conf.DBName)
	dbConn, err := sql.Open("postgres", stringConfig)
	if err != nil {
		errStr := fmt.Sprintf("[%v][BigProject][Error] : %v", time.Now(), err.Error())
		panic(errStr)
	}
	return dbConn
}
package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var Postgres *sql.DB
var err error

func ConnectDB() {

	// PostgreSqlInfo := fmt.Sprintf("%s", "port=5432 user=postgres password=king dbname=postgres")

	Postgres, err = sql.Open("postgres", "port=5432 user=postgres password=king dbname=postgres")
	if err != nil {
		log.Print("Error while connecting db :", err)
	}

	// verifying db cridentials
	if Postgres.Ping(); err != nil {
		log.Print("Incorrect Db cridentials :", err)
	} else {
		log.Print("Postgres connection success...!!!")
	}

}

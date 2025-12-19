package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Config struct {
	Username string
	Password string
	DbName   string
}

var Db *sql.DB

func OpenConnection(c Config) {
	var err error
	Db, err = sql.Open(
		"postgres",
		"user="+c.Username+
			" password="+c.Password+
			" dbname="+c.DbName+
			" sslmode=disable",
	)

	if err != nil {
		panic(err)
	}
}

func CloseConnection() {
	err := Db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

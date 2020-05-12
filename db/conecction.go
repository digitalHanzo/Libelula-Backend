package db

import (
	"database/sql"
	"fmt"

	//Repository for connect
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "libelula"
)

//Connection - connect to DB and return DB object
func Connection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s  sslmode=disable",
		host, port, user, password, dbname)

	//validation only
	db, error := sql.Open("postgres", psqlInfo)

	if error != nil {
		panic(error)
	}
	// defer db.Close()

	//connection
	error = db.Ping()
	if error != nil {
		panic(error)
	}

	println("DB succesfully connected.")
	return db
}

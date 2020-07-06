package db

import (
	"database/sql"
	"os"

	log "github.com/sirupsen/logrus"
)

//DB represents postgres DB
type DB struct {
	*sql.DB
}

//NewDB creates new postgres db
func NewDB() (*DB, error) {

	log.Debug("Opening postgres DB")

	postgresDB, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	_, err = postgresDB.Exec("CREATE TABLE IF NOT EXISTS account_info(email varchar(255),firstName varchar(40), lastName varchar(40), password varchar(40))")
	if err != nil {
		log.Fatalf("Error creating table: %q", err)
	}

	return &DB{DB: postgresDB}, nil
}

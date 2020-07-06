package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/logdna/logdna-go/logger"

	_ "github.com/lib/pq"
)

//DB represents postgres DB
type DB struct {
	*sql.DB
}

//NewDB creates new postgres db
func NewDB(log *logger.Logger) (*DB, error) {

	log.Info("Opening postgres DB")

	postgresDB, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Error(fmt.Sprintf("Error opening database: %q", err))
		return nil, err
	}

	_, err = postgresDB.Exec("CREATE TABLE IF NOT EXISTS account_info(email varchar(255), firstName varchar(40), lastName varchar(40), password varchar(40))")
	if err != nil {
		log.Error(fmt.Sprintf("Error creating table: %q", err))
		return nil, err
	}

	return &DB{postgresDB}, nil
}

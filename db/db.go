package db

import (
	"database/sql"
	"log"
	"sync"
)

type Database struct {
	Db *sql.DB
}

var instance *Database
var once sync.Once

// Connect is the "constructor" for a database. It wraps a singleton Database object so that only one connection is opened at a time
func Connect() *Database {
	once.Do(func() {
		database, err := sql.Open("sqlite3", "./sniffed.db")
		if err != nil {
			log.Fatalf("Error connecting to database: %v", err)
		}
		instance = &Database{Db: database}
	})

	return instance
}

// Query takes the question mark encoded string and all subsequent values
func (d *Database) Query(qStr string, vals ...any) (sql.Result, error) {

	tx, err := d.Db.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare(qStr)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(vals)
	if err != nil {
		return nil, err
	}
	tx.Commit()

	return res, nil
}

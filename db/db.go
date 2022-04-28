package db

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	Db *sql.DB
}

var instance *Database
var once sync.Once

// Connect is the "constructor" for a database. It wraps a singleton Database object so that only one connection is opened at a time
func Connect() *Database {
	once.Do(func() {
		// create and connect to db
		database, err := sql.Open("sqlite3", "./sniffed.db")
		if err != nil {
			log.Fatalf("Error connecting to database: %v", err)
		}
		instance = &Database{Db: database}
		_, err = instance.Db.Exec(`CREATE TABLE networks (
			uid INTEGER PRIMARY KEY AUTOINCREMENT,
			mac TEXT NOT NULL,
			power TEXT NOT NULL,
			packetCount TEXT NOT NULL,
			bssid TEXT NOT NULL,
			essid TEXT NOT NULL 
		);`)
		if err != nil {
			log.Printf("Failed to initialize database: %v", err)
		}
	})

	return instance
}

package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const dbName = "DB.db"
const dbFormat = "sqlite3"

var dbDir = "data/"
var dbPath = dbDir + dbName

// InitDB initializes the database and creates tables if they don't exist
func InitDB() error {
	// Ensure data directory exists
	if err := os.MkdirAll(dbDir, os.ModePerm); err != nil {
		fmt.Println(err)
	}
	var err error
	db := OpenDb()
	defer db.Close()
	_, err = db.Exec("PRAGMA journal_mode = WAL")
	if err != nil {
		fmt.Println(err)
	}
	// Create anime table
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS data (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		freq TEXT UNIQUE NOT NULL,
		top TEXT,
		bottom TEXT,
		left TEXT,
		right TEXT,
		middle TEXT,
		color TEXT NOT NULL,
		qr BOOLEAN
	);
	`)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func OpenDb() *sql.DB {
	db, err := sql.Open(dbFormat, dbPath)
	if err != nil {
		fmt.Println(err)
	}
	return db
}

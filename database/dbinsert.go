package database

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var err error

func Insert(freq, top, bottom, left, right, middle, color string, qr bool) {
	db := OpenDb()
	defer db.Close()

	_, err = db.Exec("INSERT OR IGNORE INTO data (freq, top, bottom, left, right, middle, color, qr) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		freq, top, bottom, left, right, middle, color, qr)
	if err != nil {
		fmt.Println(err)
	}

}

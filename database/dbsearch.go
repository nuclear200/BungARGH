package database

import (
	"fmt"
	"strings"
)

func SearchDB(query string) []string {
	db := OpenDb()
	defer db.Close()

	query = "%" + strings.TrimSpace(query) + "%" //trimming and SQL formatting

	rows, err := db.Query("SELECT name_en FROM anime WHERE name_en LIKE ?", query)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}
	var results []string
	for rows.Next() {
		var name string
		rows.Scan(&name)
		if err != nil {
			fmt.Println(err)
		}
		results = append(results, name)
	}
	if err != nil {
		fmt.Println(err)
	}
	return results
}

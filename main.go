package main

import (
	"BungieARG/database"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	database.InitDB()
	// Open the TSV file
	file, err := os.Open("data.tsv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a new CSV reader with Tab as the delimiter
	reader := csv.NewReader(file)
	reader.Comma = '\t'         // Set the delimiter to tab
	reader.FieldsPerRecord = -1 // Allow variable number of fields per record

	// Read all records
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print the records
	for _, record := range records {
		var top, bottom, left, right, center string = "", "", "", "", ""
		var qr bool
		freq := record[1]
		board := record[2]
		color := record[3]
		code, _ := strconv.Atoi(record[4])
		if code == 0 {
			qr = true
		} else {
			qr = false
		}
		perimeter := strings.Split(board, ",")
		for i := 0; i < 8; i++ {

			top = top + perimeter[i]
			bottom = bottom + perimeter[56+i]
		}
		for j := 0; j < 8; j++ {
			left = left + perimeter[j*8]
			right = right + perimeter[7+j*8]
		}
		center = perimeter[27] + perimeter[28] + perimeter[35] + perimeter[36]
		if center == "" {
			center = " "
		}
		database.Insert(freq, top, bottom, left, right, center, color, qr)
	}

}

func Extrapolate() {
	edge := "RwRwRwRwRwRwRwRw"
	database.SearchDB(edge)
}

func plot() [][][]string {
	var board [64][10][10]string
	for i := 0; i < 10; i++ {
		board[0][0][i] = "#"
		board[0][9][i] = "#"
		board[0][i][0] = "#"
		board[0][i][9] = "#"

	}
	for j := 1; j < 9; j++ {
		for l := 1; l < 9; l++ {
			board[0][j][l] = " "
		}
	}

	return board
}

func arrPrint(data [][][]string) {
	for _, board := range data {
		for _, row := range board {
			fmt.Printf("%s", strings.Join(row, ""))
		}
	}
}

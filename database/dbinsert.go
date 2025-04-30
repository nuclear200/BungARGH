package database

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var err error

func Insert(freq, top, bottom, left, right, middle string) {
	db := OpenDb()
	defer db.Close()

	_, err = db.Exec("INSERT OR IGNORE INTO data (freq, top, bottom, left, right, middle) VALUES (?, ?, ?, ?, ?, ?)",
		freq, top, bottom, left, right, middle)
	if err != nil {
		fmt.Println(err)
	}

}
func InsertBig(freq, top, bottom, left, right, middle, color string, qr bool) {
	db := OpenDb()
	defer db.Close()

	_, err = db.Exec("INSERT OR IGNORE INTO data (freq, top, bottom, left, right, middle, color, qr) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		freq, top, bottom, left, right, middle, color, qr)
	if err != nil {
		fmt.Println(err)
	}

}
func LoadData() {

	file, err := os.Open("data.json") // Open the file
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed when done

	var data map[string]string
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	for freq, board := range data {
		fmt.Println("processing: ", freq)

		switch len(string(freq)) {
		case 1:
			freq = "000" + freq
		case 2:
			freq = "00" + freq
		case 3:
			freq = "0" + freq
		default:
			continue
		}
		notation := strings.Split(board, "/")
		top := notation[0]
		bottom := notation[7]
		center := strings.Split(notation[4], "")
		var middle string
		if len(center) == 3 {
			middle = ""
		} else {
			middle = center[2]
		}

		left := calcleft(notation)
		right := calcright(notation)

		Insert(freq, top, bottom, left, right, middle)
	}
}

func calcright(notation []string) string {
	var output string = ""
	for i := 0; i < len(notation); i++ {
		str := notation[i]
		output = output + string(str[0])
	}
	return output
}

func calcleft(notation []string) string {
	var output string = ""
	len := len(notation)
	for i := 0; i < len; i++ {
		str := notation[i]
		re := regexp.MustCompile(`.$`)
		lastChar := re.FindString(str)
		output = output + lastChar
	}
	return output
}

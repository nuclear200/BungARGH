package database

import (
	"errors"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// FetchEpisodesForDownload retrieves episodes of an anime
func FetchRelations(freq string) (Top, Bottom, Left, Right string) {
	db := OpenDb()
	defer db.Close()

	// Get sides from freq
	rows, err := db.Query("SELECT top, bottom, left, right FROM data WHERE freq = ?", freq)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	i := 0
	for rows.Next() {
		i++
		err := rows.Scan(&Top, &Bottom, &Left, &Right)
		if err != nil {
			fmt.Println(err)
		}

	}
	if i == 0 {
		err = errors.New("FREQ not found!")
		fmt.Println(err)
		os.Exit(7)
	}

	return Top, Bottom, Left, Right
}

// FetchAnimeIdFromUrl retrieves the anime_id from anime table in db, given the first episode url
func FetchFreqLeft(left string) (freq string) {
	db := OpenDb()
	defer db.Close()

	err = db.QueryRow("SELECT freq FROM data WHERE right = ?", left).Scan(&freq)
	if err != nil {
		fmt.Println(err)
	}

	return freq
}

// FetchAnimeIdFromUrl retrieves the anime_id from anime table in db, given the first episode url
func FetchFreqRight(right string) (freq string) {
	db := OpenDb()
	defer db.Close()

	err = db.QueryRow("SELECT freq FROM data WHERE left = ?", right).Scan(&freq)
	if err != nil {
		fmt.Println(err)
	}

	return freq
}

// FetchAnimeIdFromUrl retrieves the anime_id from anime table in db, given the first episode url
func FetchFreqTop(top string) (freq string) {
	db := OpenDb()
	defer db.Close()

	err = db.QueryRow("SELECT freq FROM data WHERE bottom = ?", top).Scan(&freq)
	if err != nil {
		fmt.Println(err)
	}

	return freq
}

// FetchAnimeIdFromUrl retrieves the anime_id from anime table in db, given the first episode url
func FetchFreqBottom(bottom string) (freq string) {
	db := OpenDb()
	defer db.Close()

	err = db.QueryRow("SELECT freq FROM data WHERE top = ?", bottom).Scan(&freq)
	if err != nil {
		fmt.Println(err)
	}

	return freq
}

func FetchMidPiece(freq string) (center string) {
	db := OpenDb()
	defer db.Close()

	err = db.QueryRow("SELECT middle FROM data WHERE freq = ?", freq).Scan(&center)
	if err != nil {
		fmt.Println(err)
	}

	return center
}

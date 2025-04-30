package database

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// EpisodeEntry stores episode details
type EpisodeEntry struct {
	AnimeID    int
	AnimeTitle string
	Episode    int
	URL        string
}

// FetchEpisodesForDownload retrieves episodes of an anime
func FetchEpisodesForDownload(animeName string) (episodes []EpisodeEntry, DlCheck bool) {
	db := OpenDb()
	defer db.Close()

	// Get anime ID by name
	var animeID int
	var animeUrl string
	err = db.QueryRow("SELECT id, anime_url FROM anime WHERE name_en = ? OR name_jp = ?", animeName, animeName).Scan(&animeID, &animeUrl)
	if err != nil {
		fmt.Println(err)
	}

	CheckDlPresence(animeUrl)

	// Query episodes for the anime
	rows, err := db.Query("SELECT download_url, episode FROM episodes WHERE anime_id = ? AND downloaded = 0", animeID)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	// Query episodes for downloaded status
	presence, err := db.Query("SELECT downloaded FROM episodes WHERE anime_id = ?", animeID)
	if err != nil {
		fmt.Println(err)
	}
	defer presence.Close()

	var j int = 0
	//executed only for downloaded = 0
	for rows.Next() {
		var entry EpisodeEntry
		err := rows.Scan(&entry.URL, &entry.Episode)
		if err != nil {
			fmt.Println(err)
		}
		entry.AnimeID = animeID
		entry.AnimeTitle = animeName
		episodes = append(episodes, entry)
	}

	for presence.Next() {
		var val bool
		err := presence.Scan(&val)
		if err != nil {
			fmt.Println(err)
		}
		if val {
			j++
		}

	}

	if j > 0 {
		DlCheck = true
	} else {
		DlCheck = false
	}

	return episodes, DlCheck
}

// FetchAnimeIdFromUrl retrieves the anime_id from anime table in db, given the first episode url
func FetchAnimeIDFromUrl(anime_url string) (animeID int) {
	db := OpenDb()
	defer db.Close()

	err = db.QueryRow("SELECT id FROM anime WHERE anime_url = ?", anime_url).Scan(&animeID)
	if err != nil {
		fmt.Println(err)
	}

	return animeID
}

func FetchAnimeNameFromID(animeID int) (animeName string) {
	db := OpenDb()
	defer db.Close()

	err = db.QueryRow("SELECT name_en FROM anime WHERE id = ?", animeID).Scan(&animeName)
	if err != nil {
		fmt.Println(err)
	}

	return animeName
}

// retrives a string of urls from db providing anime url
func FetchEpisodesFromUrl(animeUrl string) (epUrls []string) {
	db := OpenDb()
	defer db.Close()

	// Query episodes for the anime
	rows, err := db.Query("SELECT episode_url FROM episodes WHERE episode_url LIKE ?", animeUrl)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	//executed only for downloaded = 0
	for rows.Next() {
		var epUrl string
		err := rows.Scan(&epUrl)
		if err != nil {
			fmt.Println(err)
		}

		epUrls = append(epUrls, epUrl)
	}
	return epUrls
}

func FetchEpisodesFromID(animeID int) (epUrls []string) {
	db := OpenDb()
	defer db.Close()

	// Query episodes for the anime
	rows, err := db.Query("SELECT episode_url FROM episodes WHERE anime_id = ?", animeID)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	//executed only for downloaded = 0
	for rows.Next() {
		var epUrl string
		err := rows.Scan(&epUrl)
		if err != nil {
			fmt.Println(err)
		}

		epUrls = append(epUrls, epUrl)
	}
	return epUrls
}

// ✅ Define an interface
type ScrapeInterface interface {
	ScrapeDataAndInsert(episodeUrl string)
}

// ✅ Declare global variable (default is nil)
var scraperInstance ScrapeInterface

// ✅ Function to set the scraper instance
func SetScraperInstance(s ScrapeInterface) {
	scraperInstance = s
}

func CheckDlPresence(animeUrl string) {
	if animeUrl != "" {
		db := OpenDb()
		defer db.Close()

		// Query episodes for the anime
		animeUrl = "%" + strings.TrimSpace(animeUrl) + "%"
		rows, err := db.Query("SELECT episode_url, download_url FROM episodes WHERE episode_url LIKE ?", animeUrl)
		if err != nil {
			fmt.Println(err)
		}
		defer rows.Close()

		for rows.Next() {

			var dlUrl sql.NullString
			var epUrl string
			err := rows.Scan(&epUrl, &dlUrl)
			if err != nil {
				fmt.Println(err)
			}

		}
	}

}

package main

import (
	"log"
	"sort"

	"github.com/fatih/structs"

	"github.com/korbraan/steamspy-data-extractor/database"
	"github.com/korbraan/steamspy-data-extractor/ssapi"
)

func main() {
	// Get all records
	games, err := ssapi.FetchAll()
	if err != nil {
		log.Fatal(err)
	}

	// Connect to database
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Filter records to keep new games only
	appidsInDB, err := database.GetAllAppIDs(db)
	if err != nil {
		log.Fatal(err)
	}

	gamesToInsert := []ssapi.SteamSpyGame{}
	for _, game := range games {
		sort.Ints(appidsInDB)
		// Search the current game appid in the appids from the database
		i := sort.SearchInts(appidsInDB, game.Appid)

		if !(i < len(appidsInDB) && appidsInDB[i] == game.Appid) {
			// The game is not in the database, so it needs to be added
			gamesToInsert = append(gamesToInsert, game)
		}
	}

	// Insert new games in database
	err = database.InsertBatch(db, "games", structs.Names(gamesToInsert), structs.Values(gamesToInsert))
	if err != nil {
		log.Fatal(err)
	}
}

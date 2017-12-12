package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/fatih/structs"
	"github.com/korbraan/steamspy-data-extractor/database"
	"github.com/korbraan/steamspy-data-extractor/ssapi"
	"github.com/korbraan/steamspy-data-extractor/utils"
)

func main() {
	// Get all records
	fmt.Print("Fetching all games from Steam Spy...  ")

	games, err := ssapi.FetchAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("OK")

	// Connect to database
	fmt.Print("Connecting to the database...   ")

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("OK")

	// Filter records to keep new games only
	fmt.Print("Filtering records...   ")

	appidsInDB, err := database.GetAllAppIDs(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("all app IDs gathered from database...   ")

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

	fmt.Println("OK")

	// Insert new games in database
	if len(gamesToInsert) > 0 {
		fmt.Print("Inserting " + string(len(gamesToInsert)) + " new games in database...    ")
		err = database.InsertGamesBatch(db, "games", utils.StringArrayToLower(structs.Names(gamesToInsert[0])), gamesToInsert)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("OK")
	} else {
		fmt.Println("No new game")
	}
	fmt.Println("Finished.")
}

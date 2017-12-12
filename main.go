package main

import (
	"fmt"

	"github.com/korbraan/steamspy-data-extractor/api"
	"github.com/korbraan/steamspy-data-extractor/utils"
)

func main() {
	// Get all records
	games, err := api.FetchAll()
	if err != nil {
		fmt.Print(err)
	}

	for index, game := range games {
		if utils.StringContainsCaseInsensitive(game.Name, "batman") {
			fmt.Println(string(index) + ": " + game.String())
		}
	}
}

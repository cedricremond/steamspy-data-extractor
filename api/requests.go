package api

import (
	"encoding/json"
	"net/http"
)

const steamSpyAPIRoot string = "https://steamspy.com/api.php?"

// FetchAll fetches all steam spy data and returns it as a map[userid string]SteamSpyGame
func FetchAll() (map[string]SteamSpyGame, error) {
	resp, err := http.Get(requestType("all"))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var jsonBody map[string]SteamSpyGame

	err = json.NewDecoder(resp.Body).Decode(&jsonBody)
	if err != nil {
		return nil, err
	}

	return jsonBody, nil
}

func requestType(requestType string) string {
	return steamSpyAPIRoot + "request=" + requestType
}

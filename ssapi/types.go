package ssapi

import (
	"strconv"
)

// SteamSpyGame represents a game object as fetch Steam Spy
type SteamSpyGame struct {
	Appid                  int              `json:"appid"`
	Name                   string           `json:"name"`
	Developer              string           `json:"developer"`
	Publisher              string           `json:"publisher"`
	ScoreRank              intWithStringNil `json:"score_rank"`
	Positive               int              `json:"positive"`
	Negative               int              `json:"negative"`
	UserScore              int              `json:"userscore"`
	Owners                 int              `json:"owners"`
	OwnersVariance         int              `json:"owners_variance"`
	PlayersForever         int              `json:"players_forever"`
	PlayersForeverVariance int              `json:"players_forever_variance"`
	Players2Weeks          int              `json:"players_2weeks"`
	Players2WeeksVariance  int              `json:"players_2weeks_variance"`
	AverageForever         int              `json:"average_forever"`
	Average2weeks          int              `json:"average_2weeks"`
	MedianForever          int              `json:"median_forever"`
	Median2Weeks           int              `json:"median_2weeks"`
	Price                  intWithStringNil `json:"price"`
}

func (ssp *SteamSpyGame) String() string {
	return ssp.Name + " (dev: " + ssp.Developer + ", pub: " + ssp.Publisher + ")"
}

type intWithStringNil string

func (i *intWithStringNil) UnmarshalJSON(data []byte) (err error) {
	if len(data) > 1 && data[0] == '"' && data[len(data)-1] == '"' {
		data = data[1 : len(data)-1]
	}
	if string(data) == "" || string(data) == "null" {
		data = []byte("-1")
	}
	if _, err := strconv.Atoi(string(data)); err != nil {
		return err
	}

	*i = intWithStringNil(string(data))
	return nil
}

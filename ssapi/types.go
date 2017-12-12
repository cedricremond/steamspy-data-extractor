package ssapi

// SteamSpyGame represents a game object as fetch Steam Spy
type SteamSpyGame struct {
	Appid                  int         `json:"appid"`
	Name                   string      `json:"name"`
	Developer              string      `json:"developer"`
	Publisher              string      `json:"publisher"`
	ScoreRank              interface{} `json:"score_rank"`
	Positive               int         `json:"positive"`
	Negative               int         `json:"negative"`
	Userscore              int         `json:"userscore"`
	Owners                 int         `json:"owners"`
	OwnersVariance         int         `json:"owners_variance"`
	PlayersForever         int         `json:"players_forever"`
	PlayersForeverVariance int         `json:"players_forever_variance"`
	Players2weeks          int         `json:"players_2weeks"`
	Players2weeksVariance  int         `json:"players_2weeks_variance"`
	AverageForever         int         `json:"average_forever"`
	Average2weeks          int         `json:"average_2weeks"`
	MedianForever          int         `json:"median_forever"`
	Median2weeks           int         `json:"median_2weeks"`
	Price                  string      `json:"price"`
}

func (ssp *SteamSpyGame) String() string {
	return ssp.Name + " (dev: " + ssp.Developer + ", pub: " + ssp.Publisher + ")"
}

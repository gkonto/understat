package data

type DataType int

const (
	PLAYERS = iota
	TEAMS
	GAMES
)

type Year int

type LeagueData struct {
}

type YearlyData map[int]LeagueData

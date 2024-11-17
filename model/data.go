package model

type DataType int

const (
	PLAYERS = iota
	TEAMS
	GAMES
)

type Year int

type Players string
type Teams string
type Games string

type LeagueData struct {
	Players Players
	Teams   Teams
	Games   Games
}

type YearlyData map[Year]LeagueData

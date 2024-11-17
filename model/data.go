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

type LeagueModel struct {
	Players Players
	Teams   Teams
	Games   Games
}

type LeagueData map[Year]LeagueModel

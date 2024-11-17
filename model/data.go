package model

type DataType int

const (
	PLAYERS = iota
	TEAMS
	GAMES
)

type Year int
type Players []Player
type Teams []Team
type Games []Game

type LeagueModel struct {
	Players Players
	Teams   Teams
	Games   Games
}

type LeagueData map[Year]LeagueModel

func NewLeagueModel(players Players, teams Teams, games Games) *LeagueModel {
	return &LeagueModel{
		Players: players,
		Teams:   teams,
		Games:   games,
	}
}

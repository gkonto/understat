package model

type League string

const (
	EPL        = "EPL"
	LA_LIGA    = "La_liga"
	BUNDESLIGA = "Bundesliga"
	SERIE_A    = "Serie_A"
	LIGUE_1    = "Ligue_1"
	RFPL       = "RFPL"
)

type LeagueModel struct {
	Players Players
	Teams   Teams
	Games   Games
}

type LeagueData map[Year]*LeagueModel

func NewLeagueModel(players Players, teams Teams, games Games) *LeagueModel {
	return &LeagueModel{
		Players: players,
		Teams:   teams,
		Games:   games,
	}
}

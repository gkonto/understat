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

type LeagueBundle struct {
	Page  *Page
	Model LeagueModel
}

type LeagueModel struct {
	Players *Players
	Teams   *Teams
	Games   *Games
}

type LeaguePerYear map[Year]*LeagueBundle

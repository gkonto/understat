package model

type Repository struct {
	Leagues map[League]LeagueData
}

func (p *Repository) GetLeague(league League, year Year) *LeagueModel {
	perYear, exists := p.Leagues[league]
	if exists {
		leagueModel, exists := perYear[year]
		if exists {
			return &leagueModel
		}
	}

	return nil
}

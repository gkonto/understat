package model

type Repository struct {
	Leagues map[League]LeagueData
}

func (p *Repository) GetLeague(league League, year Year) *LeagueModel {
	perYear, exists := p.Leagues[league]
	if exists {
		leagueModel, exists := perYear[year]
		if exists {
			return leagueModel
		}
	}

	return nil
}

func (p *Repository) SetModel(lmodel *LeagueModel, league League, year Year) {
	// Check if the league already exists in the Leagues map
	perYear, exists := p.Leagues[league]
	if !exists {
		p.Leagues[league] = make(LeagueData)
		perYear = p.Leagues[league]
	}

	perYear[year] = lmodel
}

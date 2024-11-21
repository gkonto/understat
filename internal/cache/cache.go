package cache

import "github.com/gkonto/understat/model"

type Repository struct {
	Leagues map[model.League]model.LeagueData
}

func NewRepository() *Repository {
	return &Repository{
		Leagues: make(map[model.League]model.LeagueData),
	}
}

func (p *Repository) GetLeague(league model.League, year model.Year) *model.LeagueModel {
	perYear, exists := p.Leagues[league]
	if exists {
		leagueModel, exists := perYear[year]
		if exists {
			return leagueModel
		}
	}

	return nil
}

func (p *Repository) SetModel(lmodel *model.LeagueModel, league model.League, year model.Year) {
	// Check if the league already exists in the Leagues map
	perYear, exists := p.Leagues[league]
	if !exists {
		p.Leagues[league] = make(model.LeagueData)
		perYear = p.Leagues[league]
	}

	perYear[year] = lmodel
}

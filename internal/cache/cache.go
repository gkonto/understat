package cache

import "github.com/gkonto/understat/model"

type Repository struct {
	Leagues map[model.League]*model.LeaguePerYear
}

func NewRepository() *Repository {
	return &Repository{
		Leagues: make(map[model.League]*model.LeaguePerYear),
	}
}

func (p *Repository) cacheLeaguesPerYear(league model.League) *model.LeaguePerYear {
	perYear, exists := p.Leagues[league]
	if !exists {
		perYear = &model.LeaguePerYear{}
		p.Leagues[league] = perYear
	}
	return perYear
}

func (p *Repository) cacheYearBundle(perYear *model.LeaguePerYear, year model.Year) *model.LeagueBundle {
	bundle, exists := (*perYear)[year]
	if !exists {
		bundle = &model.LeagueBundle{}
		(*perYear)[year] = bundle
	}
	return bundle
}

func (p *Repository) CacheBundle(league model.League, year model.Year) *model.LeagueBundle {
	perYear := p.cacheLeaguesPerYear(league)
	bundle := p.cacheYearBundle(perYear, year)
	return bundle
}

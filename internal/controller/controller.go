package controller

import (
	"github.com/gkonto/understat/internal/cache"
	"github.com/gkonto/understat/internal/requests"
	"github.com/gkonto/understat/model"
)

type RequestHandler struct {
}

type UnderstatController struct {
	Cache *cache.Repository
}

func New() *UnderstatController {
	return &UnderstatController{
		Cache: cache.NewRepository(),
	}
}

func (p *UnderstatController) GetPlayers(league model.League, year model.Year) (*model.Players, error) {
	bundle := p.Cache.CacheBundle(league, year)
	if bundle.Page == nil {
		requestHandler := requests.New()
		page, err := requestHandler.Fetch(league, year)
		if err != nil {
			return nil, err
		}
		bundle.Page = page
	}

	if bundle.Model.Players == nil {
		players, err := bundle.Page.GetPlayers()
		if err != nil {
			return nil, err
		}
		bundle.Model.Players = players
	}

	return bundle.Model.Players, nil
}

func (p *UnderstatController) GetGames(league model.League, year model.Year) (*model.Games, error) {
	bundle := p.Cache.CacheBundle(league, year)
	if bundle.Page == nil {
		requestHandler := requests.New()
		page, err := requestHandler.Fetch(league, year)
		if err != nil {
			return nil, err
		}
		bundle.Page = page
	}

	if bundle.Model.Games == nil {
		games, err := bundle.Page.GetGames()
		if err != nil {
			return nil, err
		}
		bundle.Model.Games = games
	}

	return bundle.Model.Games, nil
}

func (p *UnderstatController) GetTeams(league model.League, year model.Year) (*model.Teams, error) {
	bundle := p.Cache.CacheBundle(league, year)
	if bundle.Page == nil {
		requestHandler := requests.New()
		page, err := requestHandler.Fetch(league, year)
		if err != nil {
			return nil, err
		}
		bundle.Page = page
	}

	if bundle.Model.Teams == nil {
		teams, err := bundle.Page.GetTeams()
		if err != nil {
			return nil, err
		}
		bundle.Model.Teams = teams
	}

	return bundle.Model.Teams, nil
}

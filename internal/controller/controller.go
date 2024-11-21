package controller

import (
	"github.com/gkonto/understat/internal/cache"
	"github.com/gkonto/understat/internal/requests"
	"github.com/gkonto/understat/model"
)

type RequestHandler struct {
}

type UnderstatController struct {
	repo *cache.Repository
}

func New() *UnderstatController {
	return &UnderstatController{
		repo: cache.NewRepository(),
	}
}

func (p *UnderstatController) GetPlayers(league model.League, year model.Year) (model.Players, error) {
	leagueModel := p.repo.GetLeague(league, year)

	if leagueModel != nil {
		return leagueModel.Players, nil
	}

	leagueModel, error := p.cacheLeague(league, year)
	if error != nil {
		return nil, error
	}
	return leagueModel.Players, error
}

func (p *UnderstatController) GetGames(league model.League, year model.Year) (model.Games, error) {
	leagueModel := p.repo.GetLeague(league, year)

	if leagueModel != nil {
		return leagueModel.Games, nil
	}

	leagueModel, error := p.cacheLeague(league, year)
	if error != nil {
		return nil, error
	}
	return leagueModel.Games, error
}

func (p *UnderstatController) GetTeams(league model.League, year model.Year) (model.Teams, error) {
	leagueModel := p.repo.GetLeague(league, year)

	if leagueModel != nil {
		return leagueModel.Teams, nil
	}

	leagueModel, error := p.cacheLeague(league, year)
	if error != nil {
		return nil, error
	}
	return leagueModel.Teams, error
}

func (p *UnderstatController) cacheLeague(league model.League, year model.Year) (*model.LeagueModel, error) {
	lmodel, error := p.requestData(league, year)
	if error != nil {
		return nil, error
	}
	p.repo.SetModel(lmodel, league, year)
	return lmodel, nil
}

func (p *UnderstatController) requestData(league model.League, year model.Year) (*model.LeagueModel, error) {
	requestHandler := requests.New()
	page, err := requestHandler.Fetch(league, year)

	if err != nil {
		return nil, err
	}
	players, err := page.GetPlayers()
	if err != nil {
		return nil, err
	}
	teams, err := page.GetTeams()
	if err != nil {
		return nil, err
	}
	games, err := page.GetGames()
	if err != nil {
		return nil, err
	}

	return model.NewLeagueModel(players, teams, games), nil
}

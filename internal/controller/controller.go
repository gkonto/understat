package controller

import (
	"errors"

	"github.com/gkonto/understat/internal/requests"
	"github.com/gkonto/understat/model"
)

type UnderstatController struct {
	repo model.Repository
}

func New() *UnderstatController {
	return &UnderstatController{}
}

func (p *UnderstatController) GetPlayers(league model.League, year model.Year) (model.Players, error) {
	leagueModel := p.repo.GetLeague(league, year)

	if leagueModel != nil {
		return leagueModel.Players, nil
	}

	lmodel, error := p.requestData(league, year)
	if error != nil {
		return nil, error
	}
	p.repo.SetModel(lmodel, league, year)

	return lmodel.Players, nil
}

func (p *UnderstatController) GetGames(league model.League, year model.Year) (model.Games, error) {
	leagueModel := p.repo.GetLeague(league, year)

	if leagueModel != nil {
		return leagueModel.Games, nil
	}

	lmodel, error := p.requestData(league, year)
	if error != nil {
		return nil, error
	}
	p.repo.SetModel(lmodel, league, year)

	return lmodel.Games, nil
}

func (p *UnderstatController) GetTeams(league model.League, year model.Year) (model.Teams, error) {
	leagueModel := p.repo.GetLeague(league, year)

	if leagueModel != nil {
		return leagueModel.Teams, nil
	}

	lmodel, error := p.requestData(league, year)
	if error != nil {
		return nil, error
	}
	p.repo.SetModel(lmodel, league, year)

	return lmodel.Teams, nil
}

func (p *UnderstatController) requestData(league model.League, year model.Year) (model.LeagueModel, error) {
	requestHandler := requests.New()
	page, err := requestHandler.Fetch(league, year)

	if err != nil {
		return model.LeagueModel{}, err
	}

	lmodel, err := page.BuildModel()
	if err != nil {
		return lmodel, nil
	} else {
		return model.LeagueModel{}, errors.New("Failed to fetch LeagueModel from https://understat.com/")
	}
}

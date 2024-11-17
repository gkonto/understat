package controller

import (
	"errors"

	"github.com/gkonto/understat/model"
)

type UnderstatController struct {
	repo model.Repository
}

func New() *UnderstatController {
	return &UnderstatController{}
}

func (p *UnderstatController) GetPlayers(league model.League, year model.Year) model.Players {
	leagueModel := p.repo.GetLeague(league, year)

	if leagueModel != nil {
		return leagueModel.Players
	}

	lmodel, error := p.requestData(league, year)
	if error != nil {
		return nil
	}
	p.repo.SetModel(lmodel, league, year)

	return lmodel.Players
}

func (p *UnderstatController) GetGames(league model.League, year model.Year) model.Games {
	leagueModel := p.repo.GetLeague(league, year)

	if leagueModel != nil {
		return leagueModel.Games
	}

	lmodel, error := p.requestData(league, year)
	if error != nil {
		return nil
	}
	p.repo.SetModel(lmodel, league, year)

	return lmodel.Games
}

func (p *UnderstatController) GetTeams(league model.League, year model.Year) model.Teams {
	leagueModel := p.repo.GetLeague(league, year)

	if leagueModel != nil {
		return leagueModel.Teams
	}

	lmodel, error := p.requestData(league, year)
	if error != nil {
		return nil
	}
	p.repo.SetModel(lmodel, league, year)

	return lmodel.Teams
}

func (p *UnderstatController) requestData(league model.League, year model.Year) (model.LeagueModel, error) {
	return model.LeagueModel{}, errors.New("could not fetch LeageModel from https://understat.com/")
}

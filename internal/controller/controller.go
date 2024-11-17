package controller

import (
	"github.com/gkonto/understat/model"
)

type UnderstatController struct {
	repo model.Repository
}

func New() *UnderstatController {
	return &UnderstatController{}
}

func (p *UnderstatController) GetPlayers(league model.League, year model.Year) *model.Players {
	leagueModel := p.repo.GetLeague(league, year)

	if leagueModel == nil {
		return nil
	}

	return nil
}

func (p *UnderstatController) GetGames(league model.League, year model.Year) *model.Games {
	return nil
}

func (p *UnderstatController) GetTeams(league model.League, year model.Year) *model.Teams {
	return nil
}

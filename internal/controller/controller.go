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

func (p *UnderstatController) GetPlayers(league model.League, year int) *model.Players {
	return nil
}

func (p *UnderstatController) GetGames(league model.League, year int) *model.Games {
	return nil
}

func (p *UnderstatController) GetTeams(league model.League, year int) *model.Teams {
	return nil
}

func (p *UnderstatController) getCached(league model.League, year int) *model.LeagueData {
	return nil
}

func (p *UnderstatController) get(league model.League, year int) {

}

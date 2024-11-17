package understat

import (
	"github.com/gkonto/understat/internal/controller"
	"github.com/gkonto/understat/model"
)

type UnderstatAPI struct {
	ctrl controller.UnderstatController
}

func New() *UnderstatAPI {
	api := &UnderstatAPI{}
	return api
}

func (p *UnderstatAPI) GetPlayers(league model.League, year model.Year) (model.Players, error) {
	return p.ctrl.GetPlayers(league, year)
}

func (p *UnderstatAPI) GetGames(league model.League, year model.Year) (model.Games, error) {
	return p.ctrl.GetGames(league, year)
}

func (p *UnderstatAPI) GetTeams(league model.League, year model.Year) (model.Teams, error) {
	return p.ctrl.GetTeams(league, year)
}

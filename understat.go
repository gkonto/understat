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

// TODO this must return and error apart from model
func (p *UnderstatAPI) GetPlayers(league model.League, year model.Year) model.Players {
	return p.ctrl.GetPlayers(league, year)
}

// TODO this must return and error apart from model
func (p *UnderstatAPI) GetGames(league model.League, year model.Year) model.Games {
	return p.ctrl.GetGames(league, year)
}

// TODO this must return and error apart from model
func (p *UnderstatAPI) GetTeams(league model.League, year model.Year) model.Teams {
	return p.ctrl.GetTeams(league, year)
}

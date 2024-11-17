package understat

import (
	"github.com/gkonto/understat/internal/controller"
	"github.com/gkonto/understat/model"
)

type UnderstatAPI struct {
	ctrl controller.UnderstatController
}

func NewUnderstatAPI() *UnderstatAPI {
	api := &UnderstatAPI{}
	return api
}

func (p *UnderstatAPI) GetPlayers(league model.League, year model.Year) (model.Players, error) {
	orig, error := p.ctrl.GetPlayers(league, year)
	if error != nil {
		return nil, error
	}
	result := make(model.Players, len(orig))
	copy(result, orig)
	return result, error
}

func (p *UnderstatAPI) GetGames(league model.League, year model.Year) (model.Games, error) {
	orig, error := p.ctrl.GetGames(league, year)
	if error != nil {
		return nil, error
	}
	result := make(model.Games, len(orig))
	copy(result, orig)
	return result, error
}

func (p *UnderstatAPI) GetTeams(league model.League, year model.Year) (model.Teams, error) {
	orig, error := p.ctrl.GetTeams(league, year)
	if error != nil {
		return nil, error
	}
	result := make(model.Teams, len(orig))
	copy(result, orig)
	return result, error
}

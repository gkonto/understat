package understat

import (
	"github.com/gkonto/understat/internal/controller"
	"github.com/gkonto/understat/league"
)

type UnderstatAPI struct {
	controller controller.UnderstatController
}

func New() *UnderstatAPI {
	api := &UnderstatAPI{}
	return api
}

func (p *UnderstatAPI) GetPlayers(league league.League, year int) {
}

func (p *UnderstatAPI) GetGames(league league.League, year int) {
}

func (p *UnderstatAPI) GetTeams(league league.League, year int) {
}

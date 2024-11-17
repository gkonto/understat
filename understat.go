package understat

import (
	"github.com/gkonto/understat/data"
	"github.com/gkonto/understat/internal/controller"
	"github.com/gkonto/understat/league"
)

type UnderstatAPI struct {
	ctrl controller.UnderstatController
}

func New() *UnderstatAPI {
	api := &UnderstatAPI{}
	return api
}

func (p *UnderstatAPI) GetPlayers(league league.League, year int) {
	p.ctrl.Get(data.PLAYERS, league, year)
}

func (p *UnderstatAPI) GetGames(league league.League, year int) {
	p.ctrl.Get(data.GAMES, league, year)
}

func (p *UnderstatAPI) GetTeams(league league.League, year int) {
	p.ctrl.Get(data.TEAMS, league, year)
}

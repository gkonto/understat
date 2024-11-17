package controller

import (
	"fmt"

	"github.com/gkonto/understat/data"
	"github.com/gkonto/understat/league"
)

type UnderstatRepository struct {
	Leagues map[league.League]data.YearlyData
}

type UnderstatController struct {
	repo UnderstatRepository
}

func New() *UnderstatController {
	return &UnderstatController{}
}

func (p *UnderstatController) Get(t data.DataType, league league.League, year int) {
	switch t {
	case data.PLAYERS:
		fmt.Printf("Fetching players\n")
	case data.TEAMS:
		fmt.Printf("Fetching teams")
	case data.GAMES:
		fmt.Printf("Fetching dates")
	default:
		fmt.Printf("nil")
	}
}

func (p *UnderstatController) fetch() {

}

func (p *UnderstatController) exists() {

}

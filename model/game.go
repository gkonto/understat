package model

import "encoding/json"

type Games []Game

type Game struct {
	Id        json.Number `json:"id"`
	HasResult bool        `json:"isResult"`
	HomeTeam  Team        `json:"h"`
	AwayTeam  Team        `json:"a"`
	Goals     Goals       `json:"goals"`
	XGs       XGs         `json:"xG"`
	Datetime  string      `json:"datetime"`
	Forecast  Forecast    `json:"forecast"`
}

type Goals struct {
	Home json.Number `json:"h"`
	Away json.Number `json:"a"`
}

type XGs struct {
	Home json.Number `json:"h"`
	Away json.Number `json:"a"`
}

type Forecast struct {
	Win  json.Number `json:"w"`
	Draw json.Number `json:"d"`
	Loss json.Number `json:"l"`
}

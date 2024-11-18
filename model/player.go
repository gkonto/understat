package model

import "encoding/json"

type Player struct {
	Id        json.Number `json:"id"`
	Name      string      `json:"player_name"`
	Games     json.Number `json:"games"`
	Time      json.Number `json:"time"`
	Goals     json.Number `json:"goals"`
	XG        json.Number `json:"xG"`
	Assists   json.Number `json:"assists"`
	XA        json.Number `json:"xA"`
	Shots     json.Number `json:"shots"`
	KeyPasses json.Number `json:"key_passes"`
	Yellows   json.Number `json:"yellow_cards"`
	Reds      json.Number `json:"cards"`
	Position  string      `json:"position"`
	Team      string      `json:"team_title"`
	Npg       json.Number `json:"npg"`
	NpxG      json.Number `json:"npxG"`
	XGChain   json.Number `json:"xGChain"`
	XGBuildup json.Number `json:"xGBuildup"`
}

type Team struct {
	Id        json.Number `json:"id"`
	Name      string      `json:"title"`
	ShortName string      `json:"short_title"`
}

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

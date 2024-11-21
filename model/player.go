package model

import "encoding/json"

type Players []Player

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

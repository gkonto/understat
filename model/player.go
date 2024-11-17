package model

type Player struct {
	Id        int     `json:"id"`
	Name      string  `json:"player_name"`
	Games     int     `json:"games"`
	Time      int     `json:"time"`
	Goals     int     `json:"goals"`
	XG        float64 `json:"xG"`
	Assists   int     `json:"assists"`
	XA        float64 `json:"xA"`
	Shots     int     `json:"shots"`
	KeyPasses int     `json:"key_passes"`
	Yellows   int     `json:"yellow_cards"`
	Reds      int     `json:"cards"`
	Position  string  `json:"position"`
	Team      string  `json:"team_title"`
	Npg       int     `json:"npg"`
	NpxG      float64 `json:"npxG"`
	XGChain   float64 `json:"xGChain"`
	XGBuildup float64 `json:"xGBuildup"`
}

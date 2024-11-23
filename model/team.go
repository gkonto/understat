package model

type Teams []Team

// Define the structure for the nested objects in the JSON
type PPDA struct {
	Att int `json:"att"`
	Def int `json:"def"`
}

type History struct {
	HA          string  `json:"h_a"`
	XG          float64 `json:"xG"`
	XGA         float64 `json:"xGA"`
	NpxG        float64 `json:"npxG"`
	NpxGA       float64 `json:"npxGA"`
	PPDA        PPDA    `json:"ppda"`
	PPDAAllowed PPDA    `json:"ppda_allowed"`
	Deep        int     `json:"deep"`
	DeepAllowed int     `json:"deep_allowed"`
	Scored      int     `json:"scored"`
	Missed      int     `json:"missed"`
	Xpts        float64 `json:"xpts"`
	Result      string  `json:"result"`
	Date        string  `json:"date"`
	Wins        int     `json:"wins"`
	Draws       int     `json:"draws"`
	Loses       int     `json:"loses"`
	Pts         int     `json:"pts"`
	NpxGD       float64 `json:"npxGD"`
}

type Team struct {
	ID      string    `json:"id"`
	Title   string    `json:"title"`
	History []History `json:"history"`
}

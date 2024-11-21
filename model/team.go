package model

import "encoding/json"

type Teams []Team

type Team struct {
	Id        json.Number `json:"id"`
	Name      string      `json:"title"`
	ShortName string      `json:"short_title"`
}

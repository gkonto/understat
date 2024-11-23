package understat

import (
	"testing"

	"github.com/gkonto/understat/model"
)

// TestUnderstatAPI_GetPlayers validates fetching players for a league and year.
func TestUnderstatAPI_GetPlayers(t *testing.T) {
	api := NewUnderstatAPI()

	// Test data
	league := model.League("EPL")
	year := model.Year(2021)

	// Call the API
	players, err := api.GetPlayers(league, year)

	// Assertions
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(players) == 0 {
		t.Fatalf("expected players, got empty list")
	}
}

// TestUnderstatAPI_GetGames validates fetching games for a league and year.
func TestUnderstatAPI_GetGames(t *testing.T) {
	api := NewUnderstatAPI()

	// Test data
	league := model.League("EPL")
	year := model.Year(2021)

	// Call the API
	games, err := api.GetGames(league, year)

	// Assertions
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(games) == 0 {
		t.Fatalf("expected games, got empty list")
	}
}

// TestUnderstatAPI_GetTeams validates fetching teams for a league and year.
func TestUnderstatAPI_GetTeams(t *testing.T) {
	api := NewUnderstatAPI()

	// Test data
	league := model.League("EPL")
	year := model.Year(2021)

	// Call the API
	teams, err := api.GetTeams(league, year)

	// Assertions
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(teams) == 0 {
		t.Fatalf("expected teams, got empty list")
	}
}

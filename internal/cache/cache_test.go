package cache

import (
	"testing"

	"github.com/gkonto/understat/model"
)

func TestRepository_SetModel(t *testing.T) {
	// Create a new repository
	repo := NewRepository()

	// Prepare test data
	players := model.Players{
		{Id: "1", Name: "Player One", Goals: "5"},
		{Id: "2", Name: "Player Two", Goals: "3"},
	}

	leagueModel := &model.LeagueModel{
		Players: players,
	}

	// Add model to the repository
	repo.SetModel(leagueModel, "EPL", 2024)

	// Retrieve the model from the repository
	retrievedModel := repo.GetLeague("EPL", 2024)
	if retrievedModel == nil {
		t.Fatal("Expected model to be found in the repository, but got nil")
	}

	// Verify that the retrieved model matches the inserted model
	if len(retrievedModel.Players) != len(players) {
		t.Errorf("Expected %d players, but got %d", len(players), len(retrievedModel.Players))
	}

	if retrievedModel.Players[0].Name != "Player One" {
		t.Errorf("Expected first player name to be 'Player One', got '%s'", retrievedModel.Players[0].Name)
	}
}

func TestRepository_GetLeague_NonExistent(t *testing.T) {
	// Create a new repository
	repo := NewRepository()

	// Try to get a league model that doesn't exist
	leagueModel := repo.GetLeague("EPL", 2024)
	if leagueModel != nil {
		t.Errorf("Expected nil for non-existent league and year, but got a model")
	}
}

func TestRepository_SetModel_OverwriteExisting(t *testing.T) {
	// Create a new repository
	repo := NewRepository()

	// Prepare test data
	year := model.Year(2022)
	players := model.Players{
		{Id: "1", Name: "Player One", Goals: "5"},
	}

	// Add the first model
	leagueModel1 := &model.LeagueModel{
		Players: players,
	}
	repo.SetModel(leagueModel1, "EPL", 2024)

	// Retrieve the first model
	retrievedModel1 := repo.GetLeague("EPL", 2024)
	if retrievedModel1 == nil {
		t.Fatal("Expected model to be found in the repository, but got nil")
	}

	// Add a second model with different data (overwrite)
	players2 := model.Players{
		{Id: "2", Name: "Player Two", Goals: "3"},
	}
	leagueModel2 := &model.LeagueModel{
		Players: players2,
	}
	repo.SetModel(leagueModel2, "EPL", year)

	// Retrieve the overwritten model
	retrievedModel2 := repo.GetLeague("EPL", year)
	if retrievedModel2 == nil {
		t.Fatal("Expected model to be found after overwriting, but got nil")
	}

	// Verify that the model was correctly overwritten
	if len(retrievedModel2.Players) != len(players2) {
		t.Errorf("Expected %d players, but got %d", len(players2), len(retrievedModel2.Players))
	}

	if retrievedModel2.Players[0].Name != "Player Two" {
		t.Errorf("Expected first player name to be 'Player Two', got '%s'", retrievedModel2.Players[0].Name)
	}
}

func TestRepository_SetModel_NewLeague(t *testing.T) {
	// Create a new repository
	repo := NewRepository()

	// Prepare test data for a new league and year
	year := model.Year(2023)
	players := model.Players{
		{Id: "1", Name: "Player A", Goals: "7"},
		{Id: "2", Name: "Player B", Goals: "6"},
	}

	leagueModel := &model.LeagueModel{
		Players: players,
	}

	// Add model for new league
	repo.SetModel(leagueModel, "La_liga", year)

	// Retrieve the model for the new league
	retrievedModel := repo.GetLeague("La_liga", year)
	if retrievedModel == nil {
		t.Fatal("Expected model to be found for new league, but got nil")
	}

	// Verify that the players are correctly set for the new league
	if len(retrievedModel.Players) != len(players) {
		t.Errorf("Expected %d players, but got %d", len(players), len(retrievedModel.Players))
	}

	if retrievedModel.Players[0].Name != "Player A" {
		t.Errorf("Expected first player name to be 'Player A', got '%s'", retrievedModel.Players[0].Name)
	}
}

func TestRepository_SetModel_EmptyModel(t *testing.T) {
	// Create a new repository
	repo := NewRepository()

	// Prepare empty model
	year := model.Year(2022)
	leagueModel := &model.LeagueModel{
		Players: []model.Player{},
	}

	// Add empty model to the repository
	repo.SetModel(leagueModel, "La_liga", year)

	// Retrieve the empty model
	retrievedModel := repo.GetLeague("La_liga", year)
	if retrievedModel == nil {
		t.Fatal("Expected model to be found, but got nil")
	}

	// Verify that the model is empty
	if len(retrievedModel.Players) != 0 {
		t.Errorf("Expected 0 players, but got %d", len(retrievedModel.Players))
	}
}

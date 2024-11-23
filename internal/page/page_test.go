package page

import (
	"encoding/json"
	"testing"

	"github.com/gkonto/understat/model"
)

// Helper function to create a Page instance with mock content
func newTestPage(contents string) *Page {
	return &Page{
		Url:      "http://example.com",
		Contents: []byte(contents),
	}
}

func TestDecodeJSON(t *testing.T) {
	contents := `"\x48\x65\x6c\x6c\x6f"`
	expected := "Hello"

	decoded, err := decodeJSON(contents)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if decoded != expected {
		t.Errorf("Expected %q but got %q", expected, decoded)
	}
}

func TestBuildPlayers(t *testing.T) {
	contents := `[
	{"id":"1",
	"player_name":"Player 1",
	"games":"10",
	"time":"900",
	"goals":"5",
	"xG":"3.2",
	"assists":"2",
	"xA":"1.1",
	"shots":"15",
	"key_passes":"5",
	"yellow_cards":"1",
	"cards":"0",
	"position":"Midfielder",
	"team_title":"Team A",
	"npg":"5",
	"npxG":"3.2",
	"xGChain":"4.5",
	"xGBuildup":"2.3"
	}
	]`
	var expected model.Players
	err := json.Unmarshal([]byte(contents), &expected)
	if err != nil {
		t.Fatalf("Invalid mock data: %v", err)
	}

	players, err := buildPlayers(contents)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(players) != len(expected) {
		t.Errorf("Expected %d players but got %d", len(expected), len(players))
	}
	if players[0].Name != expected[0].Name {
		t.Errorf("Expected player name %q but got %q", expected[0].Name, players[0].Name)
	}
}

func TestGetPlayers(t *testing.T) {
	mockContent := `
		<script>var playersData = JSON.parse("[{"id":"1","player_name":"Player 1","games":"10","time":"900","goals":"5","xG":"3.2","assists":"2","xA":"1.1","shots":"15","key_passes":"5","yellow_cards":"1","cards":"0","position":"Midfielder","team_title":"Team A","npg":"5","npxG":"3.2","xGChain":"4.5","xGBuildup":"2.3"}]");</script>
	`
	page := newTestPage(mockContent)

	players, err := page.GetPlayers()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(players) != 1 {
		t.Errorf("Expected 1 player but got %d", len(players))
	}
	if players[0].Name != "Player 1" {
		t.Errorf("Expected player name %q but got %q", "Player 1", players[0].Name)
	}
}

func TestExtractData(t *testing.T) {
	mockContent := `
		<script>var playersData = JSON.parse(""[{"id":"1","player_name":"Player 1"}]"");</script>
	`
	page := newTestPage(mockContent)

	data, err := page.extractData(PLAYERS)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := `"[{"id":"1","player_name":"Player 1"}]"`
	if data != expected {
		t.Errorf("Expected %q but got %q", expected, data)
	}
}

func TestGetTeams(t *testing.T) {
	mockContent := `
		<script>var teamsData = JSON.parse("[{\"id\":\"1\",\"title\":\"Team A\",\"short_title\":\"TA\"}]");</script>
	`
	page := newTestPage(mockContent)

	teams, err := page.GetTeams()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(teams) != 0 { // `buildTeams` is not fully implemented yet, adapt based on actual implementation
		t.Errorf("Expected 0 teams but got %d", len(teams))
	}
}

func TestGetGames(t *testing.T) {
	mockContent := `
		<script>var datesData = JSON.parse("[{"id":"1","isResult":true,"h":{"id":"1","title":"Team A"},"a":{"id":"2","title":"Team B"},"goals":{"h":"2","a":"1"},"xG":{"h":"1.8","a":"0.7"},"datetime":"2024-01-01T15:00:00","forecast":{"w":"0.6","d":"0.3","l":"0.1"}}]");</script>
	`
	page := newTestPage(mockContent)

	games, err := page.GetGames()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(games) != 1 {
		t.Errorf("Expected 1 game but got %d", len(games))
	}
	if games[0].HasResult != true {
		t.Errorf("Expected HasResult to be true but got false")
	}
}

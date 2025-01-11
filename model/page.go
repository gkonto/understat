package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Page struct {
	Url      string
	Contents []byte
}

const (
	PLAYERS = "playersData"
	TEAMS   = "teamsData"
	GAMES   = "datesData"
)

func NewPage(url string, contents []byte) *Page {
	return &Page{
		Url:      url,
		Contents: contents,
	}
}

func decodeJSON(contents string) (string, error) {
	start := 0
	end := len(contents)
	if strings.HasPrefix(contents, `"`) && strings.HasSuffix(contents, `"`) {
		start = 1
		end -= 1
	}

	buffer := make([]byte, 0, len(contents))

	for i := start; i < end; i++ {
		if i+3 < end && contents[i] == '\\' && contents[i+1] == 'x' {
			// Convert \xHH to the corresponding character
			hexValue := contents[i+2 : i+4]
			var char byte
			_, err := fmt.Sscanf(hexValue, "%02X", &char)
			if err != nil {
				return "", fmt.Errorf("failed to decode hex escape: %w", err)
			}
			buffer = append(buffer, char)
			i += 3 // Skip over the escape sequence
		} else {
			buffer = append(buffer, contents[i])
		}
	}
	return string(buffer), nil
}

func buildPlayers(contents string) (*Players, error) {
	players := &Players{}
	err := json.Unmarshal([]byte(contents), players)
	if err != nil {
		return &Players{}, err
	}

	return players, nil
}

func buildTeams(contents string) (*Teams, error) {
	// Create a map to store teams by their ID
	teamsMap := map[string]Team{}

	// Unmarshal the JSON content into the map
	err := json.Unmarshal([]byte(contents), &teamsMap)
	if err != nil {
		return nil, err
	}

	// Extract the Teams from the map into a slice (array) of Team
	var teamsArray []Team
	for _, team := range teamsMap {
		teamsArray = append(teamsArray, team)
	}

	// Create a Teams object, which is a slice of Team
	teams := Teams(teamsArray)

	// Return the teams slice
	return &teams, nil
}

func buildGames(contents string) (*Games, error) {
	games := &Games{}
	err := json.Unmarshal([]byte(contents), games)
	if err != nil {
		return nil, err
	}
	return games, nil
}

func (p *Page) GetPlayers() (*Players, error) {
	playerContents, err := p.extractData(PLAYERS)
	if err != nil {
		return nil, errors.New("failed to get players data")
	}
	jsonDecoded, err := decodeJSON(playerContents)
	if err != nil {
		return nil, err
	}
	return buildPlayers(jsonDecoded)
}

func (p *Page) GetTeams() (*Teams, error) {
	teamContents, err := p.extractData(TEAMS)
	if err != nil {
		return nil, errors.New("failed to get team data")
	}
	jsonDecoded, err := decodeJSON(teamContents)
	if err != nil {
		return nil, err
	}
	return buildTeams(jsonDecoded)
}

func (p *Page) GetGames() (*Games, error) {
	gamesContents, err := p.extractData(GAMES)
	if err != nil {
		return nil, errors.New("failed to get games data")
	}
	jsonDecoded, err := decodeJSON(gamesContents)
	if err != nil {
		return nil, err
	}
	return buildGames(jsonDecoded)
}

func (p *Page) extractData(tag string) (string, error) {
	reader := strings.NewReader(string(p.Contents))
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return "", err
	}

	var jsonData string
	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		scriptContent := s.Text()
		if strings.Contains(scriptContent, tag) {
			// Extract the JSON string
			//prefix := "JSON.parse("
			start_index := p.getStartIndex(scriptContent)
			end_index := p.getEndIndex(scriptContent)

			if start_index != 0 && end_index != 0 && end_index > start_index {
				jsonData = scriptContent[start_index:end_index]
			}
		}
	})

	if jsonData == "" {
		return "", fmt.Errorf("could not parse html file contents for tag: %s", tag)
	}
	return jsonData, nil
}

func (p *Page) getStartIndex(contents string) int {
	start_index := 0
	start_pattern := `JSON\.parse\(\s*['"]`
	re_start := regexp.MustCompile(start_pattern)
	match_start := re_start.FindStringIndex(contents)
	if match_start != nil {
		start_index = match_start[1]
	}
	return start_index
}

func (p *Page) getEndIndex(contents string) int {
	end_index := 0
	end_pattern := `['"]\s*\)`
	re_end := regexp.MustCompile(end_pattern)
	match_end := re_end.FindStringIndex(contents)
	if match_end != nil {
		end_index = match_end[0]
	}
	return end_index
}

package page

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gkonto/understat/model"
)

type Page struct {
	url      string
	contents []byte
}

const (
	PLAYERS = "playersData"
	TEAMS   = "teamsData"
	GAMES   = "datesData"
)

func New(url string, contents []byte) *Page {
	return &Page{
		url:      url,
		contents: contents,
	}
}

func decodeJSON(contents string) (string, error) {
	if strings.HasPrefix(contents, `"`) && strings.HasSuffix(contents, `"`) {
		contents = contents[1 : len(contents)-1]
	}
	decoded := ""
	for i := 0; i < len(contents); i++ {
		if i+3 < len(contents) && contents[i] == '\\' && contents[i+1] == 'x' {
			// Convert \xHH to the corresponding character
			hexValue := contents[i+2 : i+4]
			var char byte
			_, err := fmt.Sscanf(hexValue, "%02X", &char)
			if err != nil {
				return "", fmt.Errorf("failed to decode hex escape: %w", err)
			}
			decoded += string(char)
			i += 3 // Skip over the escape sequence
		} else {
			decoded += string(contents[i])
		}
	}
	return decoded, nil
}

func buildPlayers(contents string) (model.Players, error) {
	fmt.Printf("%s", contents)
	return model.Players{}, nil
}

func buildTeams(contents string) (model.Teams, error) {
	return model.Teams{}, nil
}

func buildGames(contents string) (model.Games, error) {
	return model.Games{}, nil
}

func (p *Page) GetPlayers() (model.Players, error) {
	playerContents, err := p.extractData(PLAYERS)
	if err != nil {
		return nil, errors.New("Failed to get Players data")
	}
	jsonDecoded, err := decodeJSON(playerContents)
	if err != nil {
		return nil, err
	}
	return buildPlayers(jsonDecoded)
}

func (p *Page) GetTeams() (model.Teams, error) {
	teamContents, err := p.extractData(TEAMS)
	if err != nil {
		return nil, errors.New("Failed to get Team data")
	}
	jsonDecoded, err := decodeJSON(teamContents)
	if err != nil {
		return nil, err
	}
	return buildTeams(jsonDecoded)
}

func (p *Page) GetGames() (model.Games, error) {
	gamesContents, err := p.extractData(GAMES)
	if err != nil {
		return nil, errors.New("Failed to get Games data")
	}
	jsonDecoded, err := decodeJSON(gamesContents)
	if err != nil {
		return nil, err
	}
	return buildGames(jsonDecoded)
}

func (p *Page) extractData(tag string) (string, error) {
	reader := strings.NewReader(string(p.contents))
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
		return "", fmt.Errorf("Could not parse html file contents for tag: %s", tag)
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

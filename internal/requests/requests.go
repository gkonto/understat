package requests

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gkonto/understat/internal/page"
	"github.com/gkonto/understat/model"
)

type UnderstatHTMLRequest struct {
	baseURL string
	client  *http.Client
}

func New() *UnderstatHTMLRequest {
	return &UnderstatHTMLRequest{
		baseURL: "https://understat.com",
		client:  &http.Client{},
	}
}

func (p *UnderstatHTMLRequest) Fetch(league model.League, year model.Year) (*page.Page, error) {
	url := fmt.Sprintf("%s/league/%s/%d", p.baseURL, league, year)
	resp, err := p.client.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	contents, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return page.New(url, contents), nil
}

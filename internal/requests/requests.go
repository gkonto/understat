package requests

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gkonto/understat/internal/page"
	"github.com/gkonto/understat/model"
)

type HTMLGetter interface {
	Get(url string) ([]byte, error)
}

type HTTPGetter struct {
	client *http.Client
}

type UnderstatPageGetter struct {
	baseURL    string
	htmlGetter HTMLGetter
}

func New() *UnderstatPageGetter {
	return &UnderstatPageGetter{
		baseURL: "https://understat.com",
		htmlGetter: &HTTPGetter{
			client: &http.Client{},
		},
	}
}

func (p *HTTPGetter) Get(url string) ([]byte, error) {
	resp, err := p.client.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func (p *UnderstatPageGetter) FormatURL(league model.League, year model.Year) string {
	return fmt.Sprintf("%s/league/%s/%d", p.baseURL, league, year)
}

func (p *UnderstatPageGetter) Fetch(league model.League, year model.Year) (*page.Page, error) {
	url := p.FormatURL(league, year)
	contents, err := p.htmlGetter.Get(url)
	if err != nil {
		return nil, err
	}
	return page.New(url, contents), nil
}

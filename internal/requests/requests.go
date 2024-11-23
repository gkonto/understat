package requests

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gkonto/understat/model"
)

type HTMLGetter interface {
	Get(url string) ([]byte, error)
	FormatURL(league model.League, year model.Year) string
}

type HTTPGetter struct {
	client *http.Client
}

type UnderstatPageGetter struct {
	htmlGetter HTMLGetter
}

func New() *UnderstatPageGetter {
	return &UnderstatPageGetter{
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

func (p *HTTPGetter) FormatURL(league model.League, year model.Year) string {
	baseURL := "https://understat.com"
	return fmt.Sprintf("%s/league/%s/%d", baseURL, league, year)
}

func (p *UnderstatPageGetter) Fetch(league model.League, year model.Year) (*model.Page, error) {
	url := p.htmlGetter.FormatURL(league, year)
	contents, err := p.htmlGetter.Get(url)
	if err != nil {
		return nil, err
	}
	return model.NewPage(url, contents), nil
}

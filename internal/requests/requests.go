package requests

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gkonto/understat/internal/page"
	"github.com/gkonto/understat/model"
)

type HTMLGetter interface {
	Get(url string) (*page.Page, error)
}

type HTTPGetter struct {
	client *http.Client
}

type LocalPageGetter struct {
}

type UnderstatPageGetter struct {
	baseURL string
	//client     *http.Client
	htmlGetter HTMLGetter
}

func New() *UnderstatPageGetter {
	return &UnderstatPageGetter{
		baseURL: "https://understat.com",
		//client:  &http.Client{},
		htmlGetter: &HTTPGetter{
			client: &http.Client{},
		},
	}
}

func (p *HTTPGetter) Get(url string) (*page.Page, error) {
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

func (p *UnderstatPageGetter) FormatURL(league model.League, year model.Year) string {
	return fmt.Sprintf("%s/league/%s/%d", p.baseURL, league, year)
}

func (p *UnderstatPageGetter) Fetch(league model.League, year model.Year) (*page.Page, error) {
	url := p.FormatURL(league, year)
	return p.htmlGetter.Get(url)
}

package page

import "github.com/gkonto/understat/model"

type Page struct {
	url      string
	contents []byte
}

func New(url string, contents []byte) *Page {
	return &Page{
		url:      url,
		contents: contents,
	}
}

func (p *Page) BuildModel() (model.LeagueModel, error) {
	// TODO need to build the model here.
	return model.LeagueModel{}, nil
}

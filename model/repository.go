package model

type Repository struct {
	Leagues map[League]YearlyData
}

func (p *Repository) GetLeague(league League, year Year) *YearlyData {
	perYear, exists := p.Leagues[league]
	var l YearlyData
	if exists {
		l, exists = perYear[year]
		if exists {
			if l.Players != nil {
				exists = false
			}
		}
	}

	return &l
}

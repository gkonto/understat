package cache

import (
	"testing"

	"github.com/gkonto/understat/model"
)

func TestCacheLeaguesPerYear(t *testing.T) {
	repo := NewRepository()

	// Test if a new LeaguePerYear is created when league is not present
	league := model.League("Premier League")
	perYear := repo.cacheLeaguesPerYear(league)

	if perYear == nil {
		t.Errorf("Expected LeaguePerYear to be created for league %s, but it was nil", league)
	}

	if len(*perYear) != 0 {
		t.Errorf("Expected LeaguePerYear to be empty initially, but it had %d entries", len(*perYear))
	}

	// Test if it does not create a new LeaguePerYear when league already exists
	perYear2 := repo.cacheLeaguesPerYear(league)
	if perYear != perYear2 {
		t.Errorf("Expected the same LeaguePerYear for league %s, but got a different one", league)
	}
}

func TestCacheYearBundle(t *testing.T) {
	repo := NewRepository()
	league := model.League("Premier League")
	perYear := repo.cacheLeaguesPerYear(league)

	// Test if a new LeagueBundle is created when year is not present
	year := model.Year(2024)
	bundle := repo.cacheYearBundle(perYear, year)

	if bundle == nil {
		t.Errorf("Expected LeagueBundle to be created for year %d, but it was nil", year)
	}

	if len(*perYear) != 1 {
		t.Errorf("Expected PerYear to contain 1 bundle for the year, but it contained %d", len(*perYear))
	}

	if _, exists := (*perYear)[year]; !exists {
		t.Errorf("Expected PerYear to contain key %d for the year, but it was missing", year)
	}

	// Test if it does not create a new LeagueBundle when year already exists
	bundle2 := repo.cacheYearBundle(perYear, year)
	if bundle != bundle2 {
		t.Errorf("Expected the same LeagueBundle for year %d, but got a different one", year)
	}
}

func TestCacheBundle(t *testing.T) {
	repo := NewRepository()

	// Test if CacheBundle correctly caches the LeagueBundle for a given league and year
	league := model.League("La Liga")
	year := model.Year(2025)
	bundle := repo.CacheBundle(league, year)

	if bundle == nil {
		t.Errorf("Expected LeagueBundle to be returned from CacheBundle for league %s and year %d, but it was nil", league, year)
	}

	if _, exists := repo.Leagues[league]; !exists {
		t.Errorf("Expected Leagues map to contain league %s, but it was missing", league)
	}

	perYear, exists := repo.Leagues[league]
	if !exists {
		t.Errorf("Expected PerYear to exist for league %s, but it was missing", league)
	}

	// Assert that the bundle is cached correctly in the perYear map
	if _, exists := (*perYear)[year]; !exists {
		t.Errorf("Expected PerYear to contain the year %d, but it was missing", year)
	}
}

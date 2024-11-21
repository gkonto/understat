package requests

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/gkonto/understat/model"
)

type MockHTMLGetter struct {
}

type MockHTMLGetterError struct {
}

func (p *MockHTMLGetterError) Get(url string) ([]byte, error) {
	return nil, fmt.Errorf("An error occured")
}

func (p *MockHTMLGetterError) FormatURL(league model.League, year model.Year) string {
	return ""
}

func (p *MockHTMLGetterError) String() string {
	return "An error occured"
}

func (p *MockHTMLGetter) Get(url string) ([]byte, error) {
	return []byte("hello friend"), nil
}

func (p *MockHTMLGetter) FormatURL(league model.League, year model.Year) string {
	baseURL := "https://understat.com"
	return fmt.Sprintf("%s/league/%s/%d", baseURL, league, year)
}

func TestUnderstatPageGetter_FormatURL(t *testing.T) {
	// Create a new UnderstatPageGetter
	understatGetter := HTTPGetter{}

	// Test with a sample league and year
	league := "Bundesliga" // replace with an actual value from model.League
	year := model.Year(2023)

	// Call FormatURL and check the result
	expectedURL := fmt.Sprintf("https://understat.com/league/%s/%d", league, year)
	result := understatGetter.FormatURL("Bundesliga", year)

	if expectedURL != result {
		t.Errorf("The formatted URL should match the expected URL.")
	}
}

func TestUnderstatPageGetter_Fetch_Success(t *testing.T) {
	// Create a mock HTMLGetter and set up expected behavior
	mockGetter := new(MockHTMLGetter)
	understatGetter := &UnderstatPageGetter{
		htmlGetter: mockGetter,
	}

	// Prepare the mock data
	year := model.Year(2023)
	expected := []byte("hello friend")

	// Call Fetch and verify the result
	page, err := understatGetter.Fetch("Bundesliga", year)
	if err != nil {
		t.Errorf("Should expect no error here: Got: %s", err)
	}
	if !bytes.Equal(expected, page.Contents) {
		t.Errorf("Expected: %s Got: %s", expected, page.Contents)
	}
}

func TestUnderstatPageGetter_Fetch_Error(t *testing.T) {
	// Create a mock HTMLGetter and set up expected behavior
	mockGetter := new(MockHTMLGetterError)
	understatGetter := &UnderstatPageGetter{
		htmlGetter: mockGetter,
	}

	// Prepare the mock data
	year := model.Year(2023)

	// Call Fetch and verify the result
	page, err := understatGetter.Fetch("Bundesliga", year)

	if page != nil {
		t.Errorf("No page should be returned, Got: %s", page.Contents)
	}

	// Assert an error occurred
	if err.Error() != "An error occured" {
		t.Errorf("Expected error: %s, Got: %s", mockGetter.String(), err)
	}
}

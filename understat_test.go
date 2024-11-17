package understat

import (
	"testing"

	"github.com/gkonto/understat/model"
)

func TestGetPlayers(t *testing.T) {
	api := NewUnderstatAPI()
	api.GetPlayers(model.EPL, 2024)
}

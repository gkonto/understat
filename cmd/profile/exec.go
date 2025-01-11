package main

import (
	"fmt"
	"os"
	"runtime/pprof"

	"github.com/gkonto/understat"
	"github.com/gkonto/understat/model"
)

func main() {

	f, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Println("Could not create CPU profile:", err)
		return
	}

	defer f.Close()

	if err := pprof.StartCPUProfile(f); err != nil {
		fmt.Println("Could not start CPU profile:", err)
		return
	}
	defer pprof.StopCPUProfile()

	api := understat.NewUnderstatAPI()
	teams, err := api.GetTeams(model.League("EPL"), model.Year(2023))

	if err != nil {
		fmt.Println("Error fetching teams:", err)
		return
	}
	fmt.Println("Teams:", len(teams))
}

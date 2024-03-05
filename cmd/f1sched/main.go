package main

import (
	"log"

	"github.com/lucas-ingemar/f1sched/internal/api"
	"github.com/lucas-ingemar/f1sched/internal/f1com"
	"github.com/lucas-ingemar/f1sched/internal/tui"
)

func main() {

	raceApi := api.NewRaceApi(f1com.F1com{})
	raceSchedule, err := raceApi.RaceSchedule(2024)
	if err != nil {
		log.Fatal(err)
	}

	// races, err = ergast.GetRaceData()
	// if err != nil {
	// 	panic(err)
	// }

	if err := tui.Run(raceSchedule); err != nil {
		log.Fatal(err)
	}
}

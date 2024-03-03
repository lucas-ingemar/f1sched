package main

import (
	"log"

	"github.com/lucas-ingemar/f1sched/internal/ergast"
	"github.com/lucas-ingemar/f1sched/internal/tui"
)

func main() {
	races, err := ergast.GetRaceData()
	if err != nil {
		panic(err)
	}

	if err := tui.Run(races); err != nil {
		log.Fatal(err)
	}
}

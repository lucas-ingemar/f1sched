package ergast

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/lucas-ingemar/f1sched/internal/shared"
)

const (
	SCHEDULE_URL = "http://ergast.com/api/f1/2024.json"
)

func GetRaceData() (races []shared.Race, err error) {
	data, err := downloadRaceData()
	if err != nil {
		return nil, err
	}

	for idx, r := range data.RaceTable.Races {
		_ = idx
		// if idx > 6 {
		// 	continue
		// }
		nr, err := generateRace(r)
		if err != nil {
			return nil, err
		}
		races = append(races, nr)
	}

	return
}

func generateRace(eRace shared.ErgastRace) (shared.Race, error) {
	rtV, err := time.Parse("2006-01-02T15:04:05Z", fmt.Sprintf("%sT%s", eRace.Date, eRace.Time))
	rt := &rtV
	if err != nil {
		rtV, _ = time.Parse("2006-01-02", eRace.Date)
		rtV = rtV.In(time.Local)
		rt = &rtV
	}

	r := shared.Race{
		Round:          eRace.Round,
		RaceName:       eRace.RaceName,
		Circuit:        eRace.Circuit.CircuitName,
		Country:        eRace.Circuit.Location.Country,
		FirstPractice:  parseTime(eRace.FirstPractice.Date, eRace.FirstPractice.Time),
		SecondPractice: parseTime(eRace.SecondPractice.Date, eRace.SecondPractice.Time),
		Qualifying:     parseTime(eRace.Qualifying.Date, eRace.Qualifying.Time),
		Race:           rt,
	}

	if eRace.Sprint != nil {
		r.Type = shared.SprintRace
		r.Sprint = parseTime(eRace.Sprint.Date, eRace.Sprint.Time)
	} else {
		if eRace.ThirdPractice != nil {
			r.ThirdPractice = parseTime(eRace.ThirdPractice.Date, eRace.ThirdPractice.Time)
		}
		r.Type = shared.NormalRace
	}

	return r, nil
}

func parseTime(d, t string) *time.Time {
	tV, err := time.Parse("2006-01-02T15:04:05Z", fmt.Sprintf("%sT%s", d, t))
	if err != nil {
		return nil
	}
	nt := tV.In(time.Local)
	return &nt
}

func downloadRaceData() (shared.MRData, error) {
	resp, err := http.Get(SCHEDULE_URL)
	if err != nil {
		return shared.MRData{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return shared.MRData{}, err
	}

	var data shared.Ergast
	err = json.Unmarshal(body, &data)
	if err != nil {
		return shared.MRData{}, err
	}

	return data.MRData, nil
}

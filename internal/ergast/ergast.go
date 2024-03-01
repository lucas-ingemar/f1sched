package ergast

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/lucas-ingemar/f1sched/internal/shared"
)

func GetRaceData() (races []shared.Race, err error) {
	data, err := downloadRaceData()
	if err != nil {
		return nil, err
	}

	for _, r := range data.RaceTable.Races {
		nr, err := generateRace(r)
		if err != nil {
			return nil, err
		}
		races = append(races, nr)
	}

	return
}

func generateRace(eRace shared.ErgastRace) (shared.Race, error) {
	fp1, err := time.Parse("2006-01-02T15:04:05Z", fmt.Sprintf("%sT%s", eRace.FirstPractice.Date, eRace.FirstPractice.Time))
	if err != nil {
		return shared.Race{}, err
	}

	fp2, err := time.Parse("2006-01-02T15:04:05Z", fmt.Sprintf("%sT%s", eRace.SecondPractice.Date, eRace.SecondPractice.Time))
	if err != nil {
		return shared.Race{}, err
	}

	rt, err := time.Parse("2006-01-02T15:04:05Z", fmt.Sprintf("%sT%s", eRace.Date, eRace.Time))
	if err != nil {
		return shared.Race{}, err
	}

	qt, err := time.Parse("2006-01-02T15:04:05Z", fmt.Sprintf("%sT%s", eRace.Qualifying.Date, eRace.Qualifying.Time))
	if err != nil {
		return shared.Race{}, err
	}

	r := shared.Race{
		Round:          eRace.Round,
		RaceName:       eRace.RaceName,
		Circuit:        eRace.Circuit.CircuitName,
		Country:        eRace.Circuit.Location.Country,
		FirstPractice:  fp1,
		SecondPractice: fp2,
		Qualifying:     qt,
		Race:           rt,
	}

	if eRace.Sprint != nil {
		st, err := time.Parse("2006-01-02T15:04:05Z", fmt.Sprintf("%sT%s", eRace.Sprint.Date, eRace.Sprint.Time))
		if err != nil {
			return shared.Race{}, err
		}
		r.Type = shared.SprintRace
		r.Sprint = st
	} else {
		fp3, err := time.Parse("2006-01-02T15:04:05Z", fmt.Sprintf("%sT%s", eRace.ThirdPractice.Date, eRace.ThirdPractice.Time))
		if err != nil {
			return shared.Race{}, err
		}
		r.Type = shared.NormalRace
		r.ThirdPractice = fp3
	}

	return r, nil
}

func downloadRaceData() (shared.MRData, error) {
	resp, err := http.Get("http://ergast.com/api/f1/current.json")
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

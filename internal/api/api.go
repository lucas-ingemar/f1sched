package api

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/lucas-ingemar/f1sched/internal/shared"
)

type RaceApiFace interface {
	RaceScheduleRaw(year int) (shared.RaceSchedule, error)
}

type RaceApi struct {
	RaceApiFace
}

func (ra RaceApi) RaceSchedule(year int) (rs shared.RaceSchedule, err error) {
	cacheDir := filepath.Join(xdg.CacheHome, "f1sched")
	jsonFilename := filepath.Join(cacheDir, fmt.Sprintf("%d_race_schedule.json", year))

	err = os.MkdirAll(cacheDir, os.ModePerm)
	if err != nil {
		if !os.IsExist(err) {
			return shared.RaceSchedule{}, err
		}
	}

	jsonInputData, err := os.ReadFile(jsonFilename)
	if err == nil {
		err = json.Unmarshal(jsonInputData, &rs)
		if err == nil && rs.FullInformation {
			return
		}
	}

	rs, err = ra.RaceScheduleRaw(year)
	if err != nil {
		return shared.RaceSchedule{}, err
	}

	jsonOutputData, err := json.MarshalIndent(rs, "", "    ")
	if err != nil {
		return rs, err
	}

	err = os.WriteFile(jsonFilename, jsonOutputData, 0644)
	if err != nil {
		return rs, err
	}

	return rs, nil
}

func NewRaceApi(raceApi RaceApiFace) RaceApi {
	return RaceApi{
		RaceApiFace: raceApi,
	}
}

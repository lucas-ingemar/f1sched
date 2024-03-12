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
	DriverStandingsRaw(rs shared.RaceSchedule) (ds shared.DriverStandings, err error)
}

type RaceApi struct {
	RaceApiFace
}

func (ra RaceApi) RaceSchedule(year int) (rs shared.RaceSchedule, err error) {
	cacheDir := filepath.Join(xdg.CacheHome, "f1sched", fmt.Sprint(year))
	jsonFilename := filepath.Join(cacheDir, "race_schedule.json")

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

func (ra RaceApi) DriverStandings(rs shared.RaceSchedule) (ds shared.DriverStandings, err error) {
	cacheDir := filepath.Join(xdg.CacheHome, "f1sched", fmt.Sprint(rs.Year))
	jsonFilename := filepath.Join(cacheDir, "driver_standings.json")

	err = os.MkdirAll(cacheDir, os.ModePerm)
	if err != nil {
		if !os.IsExist(err) {
			return ds, err
		}
	}

	jsonInputData, err := os.ReadFile(jsonFilename)
	if err == nil {
		err = json.Unmarshal(jsonInputData, &ds)
		lr := rs.GetLatestFinishedRace()
		if err == nil && lr != nil && lr.EndTime.Before(ds.UpdatedAt) {
			return
		}
	}

	ds, err = ra.DriverStandingsRaw(rs)
	if err != nil {
		return shared.DriverStandings{}, err
	}

	jsonOutputData, err := json.MarshalIndent(ds, "", "    ")
	if err != nil {
		return ds, err
	}

	err = os.WriteFile(jsonFilename, jsonOutputData, 0644)
	if err != nil {
		return ds, err
	}

	return ds, nil
}

func NewRaceApi(raceApi RaceApiFace) RaceApi {
	return RaceApi{
		RaceApiFace: raceApi,
	}
}

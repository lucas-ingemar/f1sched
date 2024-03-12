package shared

import "time"

type RaceType string

const (
	NormalRace RaceType = "normal"
	SprintRace RaceType = "sprint"
)

type RaceSchedule struct {
	Year            int               `json:"year"`
	FullInformation bool              `json:"full_information"`
	Races           []RaceInformation `json:"races"`
}

func (rs RaceSchedule) GetLatestFinishedRace() *RaceInformation {
	for idx, r := range rs.Races {
		if r.EndTime.After(time.Now()) && idx > 0 {
			return &rs.Races[idx-1]
		} else if r.EndTime.After(time.Now()) && idx == 0 {
			return nil
		}
	}
	return &rs.Races[len(rs.Races)-1]
}

type RaceInformation struct {
	Name             string       `json:"name"`
	Type             RaceType     `json:"type"`
	Location         RaceLocation `json:"location"`
	Round            int          `json:"round"`
	StartTime        time.Time    `json:"start_time"`
	EndTime          time.Time    `json:"end_time"`
	FreePractice1    *RaceEvent   `json:"free_practice_1,omitempty"`
	FreePractice2    *RaceEvent   `json:"free_practice_2,omitempty"`
	FreePractice3    *RaceEvent   `json:"free_practice_3,omitempty"`
	SprintQualifying *RaceEvent   `json:"sprint_qualifying,omitempty"`
	Sprint           *RaceEvent   `json:"sprint,omitempty"`
	Qualifying       *RaceEvent   `json:"qualifying,omitempty"`
	Race             *RaceEvent   `json:"race,omitempty"`
}

type RaceLocation struct {
	Country string `json:"country"`
	City    string `json:"city"`
	Circuit string `json:"circuit"`
}

type RaceEvent struct {
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

type DriverStandings struct {
	UpdatedAt time.Time        `json:"updated_at"`
	Entries   []StandingDriver `json:"entries"`
}

type StandingDriver struct {
	Position        int    `json:"position"`
	DriverFirstName string `json:"driver_first_name"`
	DriverLastName  string `json:"driver_last_name"`
	DriverShort     string `json:"driver_short"`
	Nationality     string `json:"nationality"`
	Team            string `json:"team"`
	Points          int    `json:"points"`
}

//////////////////////////////////////

type Race struct {
	Round          string
	RaceName       string
	Circuit        string
	Country        string
	Type           RaceType
	FirstPractice  *time.Time
	SecondPractice *time.Time
	ThirdPractice  *time.Time
	Qualifying     *time.Time
	Sprint         *time.Time
	Race           *time.Time
}

type Location struct {
	Lat      string `json:"lat"`
	Long     string `json:"long"`
	Locality string `json:"locality"`
	Country  string `json:"country"`
}

type Circuit struct {
	CircuitId   string   `json:"circuitId"`
	URL         string   `json:"url"`
	CircuitName string   `json:"circuitName"`
	Location    Location `json:"Location"`
}

type Practice struct {
	Date string `json:"date"`
	Time string `json:"time"`
}

type ErgastRace struct {
	Season         string    `json:"season"`
	Round          string    `json:"round"`
	URL            string    `json:"url"`
	RaceName       string    `json:"raceName"`
	Circuit        Circuit   `json:"Circuit"`
	Date           string    `json:"date"`
	Time           string    `json:"time"`
	FirstPractice  Practice  `json:"FirstPractice"`
	SecondPractice Practice  `json:"SecondPractice"`
	Qualifying     Practice  `json:"Qualifying"`
	Sprint         *Practice `json:"Sprint,omitempty"`
	ThirdPractice  *Practice `json:"ThirdPractice,omitempty"`
}

type MRData struct {
	XMLNS     string `json:"xmlns"`
	Series    string `json:"series"`
	URL       string `json:"url"`
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	Total     string `json:"total"`
	RaceTable struct {
		Season string       `json:"season"`
		Races  []ErgastRace `json:"Races"`
	} `json:"RaceTable"`
}

type Ergast struct {
	MRData MRData `json:"MRData"`
}

type Color struct {
	Color1 string
	Color2 string
	Color3 string
}

type Country struct {
	Name    string
	BgColor Color
	FgColor Color
	Flag    string
}

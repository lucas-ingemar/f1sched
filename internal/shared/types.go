package shared

import "time"

type RaceType string

const (
	NormalRace RaceType = "normal"
	SprintRace RaceType = "sprint"
)

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
}

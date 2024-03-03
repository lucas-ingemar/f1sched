package tui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lucas-ingemar/f1sched/internal/countries"
	"github.com/lucas-ingemar/f1sched/internal/shared"
)

var (
	normalStyle = lipgloss.NewStyle()
	pastStyle   = lipgloss.NewStyle().Foreground(lipgloss.ANSIColor(8))
)

type RaceSummaryModel struct {
	Style lipgloss.Style

	race    shared.Race
	country shared.Country

	frame int
	id    int
	tag   int
}

func (m RaceSummaryModel) Init() tea.Cmd {
	return nil
}
func (m RaceSummaryModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m RaceSummaryModel) View() string {
	s := ""
	s += lipgloss.NewStyle().
		Background(lipgloss.Color(m.country.BgColor.Color1)).
		Foreground(lipgloss.Color(m.country.FgColor.Color1)).
		Align(lipgloss.Center).
		Width(34).
		Render(m.race.Country)
	s += "\n      "
	s += lipgloss.NewStyle().
		Background(lipgloss.Color(m.country.BgColor.Color2)).
		Foreground(lipgloss.Color(m.country.FgColor.Color2)).
		Align(lipgloss.Center).
		Width(34).
		Render(m.race.Circuit)
	s += "\n   "
	s += lipgloss.NewStyle().
		Background(lipgloss.Color(m.country.BgColor.Color3)).
		Foreground(lipgloss.Color(m.country.FgColor.Color3)).
		Align(lipgloss.Center).
		Width(34).
		Render(m.race.RaceName)
	s += "\n"
	s += "\n"
	s += m.race.Race.AddDate(0, 0, -2).In(time.Local).Format(" Jan 02 - ") + m.race.Race.In(time.Local).Format("Jan 02, 2006")
	if m.race.Type == shared.SprintRace {
		s += "         "
		s += lipgloss.NewStyle().Background(lipgloss.ANSIColor(8)).Render(" Sprint ")
	}
	s += "\n"
	s += "\n"

	// s += lipgloss.NewStyle().
	// 	Align(lipgloss.Right).
	// 	Width(30).
	// 	Render("Practice 1: 29 Feb, 15:00")
	// 	 2006-01-02T15:04:05Z
	s += getDateStyle(m.race.FirstPractice).Render("         Practice 1:      " + formatDateTime(m.race.FirstPractice))
	s += "\n"
	s += getDateStyle(m.race.SecondPractice).Render("         Practice 2:      " + formatDateTime(m.race.SecondPractice))
	s += "\n"
	if m.race.Type == shared.NormalRace {
		s += getDateStyle(m.race.ThirdPractice).Render("         Practice 3:      " + formatDateTime(m.race.ThirdPractice))
		s += "\n"
		s += "\n"
		s += getDateStyle(m.race.Qualifying).Render("         Qualifying:      " + formatDateTime(m.race.Qualifying))
	} else {
		s += "\n"
		s += getDateStyle(m.race.Qualifying).Render("         Qualifying:      " + formatDateTime(m.race.Qualifying))
		s += "\n"
		s += getDateStyle(m.race.Sprint).Render("         Sprint:          " + formatDateTime(m.race.Sprint))
	}
	s += "\n"
	s += getDateStyle(m.race.Race).Render("         Race:            " + formatDateTime(m.race.Race))
	// s += "         Race:            02 Mar, 15:00"
	// s += "         Race:           Max Verstappen"

	return s
}

// New returns a model with default values.
func NewRaceSummary(race shared.Race) RaceSummaryModel {
	m := RaceSummaryModel{
		race:    race,
		country: countries.GetCountry(race.Country),
	}

	return m
}

func formatDateTime(dt *time.Time) string {
	if dt == nil {
		return "*** **, **:**"
	}
	return dt.In(time.Local).Format("Jan 02, 15:04")
}

func getDateStyle(t *time.Time) lipgloss.Style {
	if t == nil {
		return normalStyle
	}

	if t.Add(2 * time.Hour).Before(time.Now()) {
		return pastStyle
	}
	return normalStyle
}

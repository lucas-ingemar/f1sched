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

	race    shared.RaceInformation
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
		Render(m.race.Location.Country)
	s += "\n      "
	s += lipgloss.NewStyle().
		Background(lipgloss.Color(m.country.BgColor.Color2)).
		Foreground(lipgloss.Color(m.country.FgColor.Color2)).
		Align(lipgloss.Center).
		Width(34).
		Render(m.race.Location.Circuit)
	s += "\n   "
	s += lipgloss.NewStyle().
		Background(lipgloss.Color(m.country.BgColor.Color3)).
		Foreground(lipgloss.Color(m.country.FgColor.Color3)).
		Align(lipgloss.Center).
		Width(34).
		Render(m.race.Name)
	s += "\n"
	s += "\n"
	s += m.race.StartTime.In(time.Local).Format(" Jan 02 - ") + m.race.EndTime.In(time.Local).Format("Jan 02, 2006")
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
	s += getDateStyle(m.race.FreePractice1).Render("         Practice 1:      " + formatDateTime(&m.race.FreePractice1.StartTime))
	s += "\n"
	if m.race.Type == shared.NormalRace {
		s += getDateStyle(m.race.FreePractice2).Render("         Practice 2:      " + formatDateTime(&m.race.FreePractice2.StartTime))
		s += "\n"
		s += getDateStyle(m.race.FreePractice3).Render("         Practice 3:      " + formatDateTime(&m.race.FreePractice3.StartTime))
		s += "\n"
		s += "\n"
		s += getDateStyle(m.race.Qualifying).Render("         Qualifying:      " + formatDateTime(&m.race.Qualifying.StartTime))
	} else {
		s += "\n"
		s += getDateStyle(m.race.SprintQualifying).Render("  Sprint Qualifying:      " + formatDateTime(&m.race.SprintQualifying.StartTime))
		s += "\n"
		s += getDateStyle(m.race.Sprint).Render("         Sprint:          " + formatDateTime(&m.race.Sprint.StartTime))
		s += "\n"
		s += getDateStyle(m.race.Qualifying).Render("         Qualifying:      " + formatDateTime(&m.race.Qualifying.StartTime))
	}
	s += "\n"
	s += getDateStyle(m.race.Race).Render("         Race:            " + formatDateTime(&m.race.Race.StartTime))
	// s += "         Race:            02 Mar, 15:00"
	// s += "         Race:           Max Verstappen"

	return s
}

// New returns a model with default values.
func NewRaceSummary(race shared.RaceInformation) RaceSummaryModel {
	m := RaceSummaryModel{
		race:    race,
		country: countries.GetCountry(race.Location.Country),
	}

	return m
}

func formatDateTime(dt *time.Time) string {
	if dt == nil {
		return "*** **, **:**"
	}
	return dt.In(time.Local).Format("Jan 02, 15:04")
}

func getDateStyle(event *shared.RaceEvent) lipgloss.Style {
	if event == nil {
		return normalStyle
	}

	if event.EndTime.Before(time.Now()) {
		return pastStyle
	}
	return normalStyle
}

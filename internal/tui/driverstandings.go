package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lucas-ingemar/f1sched/internal/countries"
	"github.com/lucas-ingemar/f1sched/internal/shared"
)

// type sessionState uint

var (
	dsTableHeader        = []string{"Pos", "Nationality", "Driver", "Team", "Pts"}
	dsTableHeaderStyle   = lipgloss.NewStyle().Bold(true).BorderForeground(lipgloss.ANSIColor(8)).Border(lipgloss.NormalBorder(), false, false, true, false)
	dsTableLineStyleEven = lipgloss.NewStyle().Background(lipgloss.ANSIColor(8))
	dsTableLineStyleOdd  = lipgloss.NewStyle()
)

type driverStandingsModel struct {
	driverStandings shared.DriverStandings
	// races          []RaceSummaryModel
	// selectedIndex  int
	// row            int
	// col            int
	terminalWidth  int
	terminalHeight int
	// cardsPerRow    int
	// totalRows      int
}

func newDriverStandingsModel(driverStandings shared.DriverStandings) driverStandingsModel {
	m := driverStandingsModel{
		driverStandings: driverStandings,
	}
	// for _, r := range raceSchedule.Races {
	// 	m.races = append(m.races, NewRaceSummary(r))
	// }
	return m
}

func (m driverStandingsModel) Init() tea.Cmd {
	models := []tea.Cmd{}
	// for _, t := range m.races {
	// 	models = append(models, t.Init())
	// }
	return tea.Batch(models...)
}

func (m driverStandingsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.terminalWidth = msg.Width
		m.terminalHeight = msg.Height
		// 	m.cardsPerRow = int(m.terminalWidth / cardWidth)
		// 	m.totalRows = int(math.Ceil(float64(len(m.races)) / float64(m.cardsPerRow)))
		// case tea.KeyMsg:
		// 	switch msg.String() {
		// 	case "up", "down", "left", "right":
		// 		err := m.moveFocus(msg.String())
		// 		_ = err
		// 	case "ctrl+c", "q":
		// 		return m, tea.Quit
		// 		// case "tab":
		// 		// 	if m.state == timerView {
		// 		// 		m.state = spinnerView
		// 		// 	} else {
		// 		// 		m.state = timerView
		// 		// 	}
		// 		// case "n":
		// 		// 	if m.state == timerView {
		// 		// 		m.timer = timer.New(defaultTime)
		// 		// 		cmds = append(cmds, m.timer.Init())
		// 		// 	} else {
		// 		// 		// m.Next()
		// 		// 		// m.resetSpinner()
		// 		// 		cmds = append(cmds, m.spinner.Tick)
		// 		// 	}
		// 	}
		// 	// switch m.state {
		// 	// // update whichever model is focused
		// 	// case spinnerView:
		// 	// 	m.spinner, cmd = m.spinner.Update(msg)
		// 	// 	cmds = append(cmds, cmd)
		// 	// default:
		// 	// 	m.timer, cmd = m.timer.Update(msg)
		// 	// 	cmds = append(cmds, cmd)
		// 	// }
		// 	// case spinner.TickMsg:
		// 	// 	m.spinner, cmd = m.spinner.Update(msg)
		// 	// 	cmds = append(cmds, cmd)
		// 	// case timer.TickMsg:
		// 	// 	m.timer, cmd = m.timer.Update(msg)
		// 	// 	cmds = append(cmds, cmd)
	}
	return m, tea.Batch(cmds...)
}

func (m driverStandingsModel) View() string {
	doc := strings.Builder{}

	lines := []string{
		// dsTableHeaderStyle.Width(m.terminalWidth).Render(strings.Join(dsTableHeader, " ")),
	}

	cSizes := m.columnSizes()

	headerCols := []string{}
	for idx, dts := range dsTableHeader {
		headerCols = append(headerCols, dsTableHeaderStyle.Width(cSizes[idx]).Render(dts))
	}

	lines = append(lines, lipgloss.JoinHorizontal(lipgloss.Left, headerCols...))

	for idx, ds := range m.driverStandings.Entries {
		var style lipgloss.Style
		if idx%2 == 0 {
			style = dsTableLineStyleEven
		} else {
			style = dsTableLineStyleOdd
		}

		style = dsTableLineStyleOdd
		cols := []string{}
		cols = append(cols, style.Width(cSizes[0]).Render(fmt.Sprint(ds.Position)))
		cols = append(cols, style.Width(cSizes[1]).Render(generateFlag(ds.Nationality)))

		var col3text string
		driverFullname := fmt.Sprintf("%s %s", ds.DriverFirstName, ds.DriverLastName)
		if cSizes[2]-1 > len(driverFullname) {
			col3text = driverFullname
		} else if cSizes[2]-1 > len(ds.DriverLastName) {
			col3text = ds.DriverLastName
		} else {
			col3text = ds.DriverShort
		}

		cols = append(cols, style.Width(cSizes[2]).Render(col3text))
		cols = append(cols, style.Width(cSizes[3]).Render(getTeamStyle(ds.Team).Render(ds.Team)))
		cols = append(cols, style.Width(cSizes[4]).Render(fmt.Sprint(ds.Points)))

		lines = append(lines, lipgloss.JoinHorizontal(lipgloss.Left, cols...))
	}
	lines = append(lines, "")
	doc.WriteString(lipgloss.JoinVertical(lipgloss.Top, lines...))
	return doc.String()
}

func (m driverStandingsModel) columnSizes() []int {
	c1 := []string{dsTableHeader[0]}
	c2 := []string{dsTableHeader[1]}
	c4 := []string{dsTableHeader[3]}
	c5 := []string{dsTableHeader[4]}
	for _, d := range m.driverStandings.Entries {
		c1 = append(c1, fmt.Sprint(d.Position))
		c2 = append(c2, d.Nationality)
		c4 = append(c4, d.Team)
		c5 = append(c5, fmt.Sprint(d.Points))
	}

	c1s := lenLongestStr(c1) + 2
	c2s := lenLongestStr(c2) + 2
	c4s := lenLongestStr(c4) + 2
	c5s := lenLongestStr(c5) + 2

	c3s := m.terminalWidth - (c1s + c2s + c4s + c5s)

	return []int{c1s, c2s, c3s, c4s, c5s}
}

func lenLongestStr(strings []string) (length int) {
	for _, s := range strings {
		length = max(length, len(s))
	}
	return
}

func generateFlag(country string) (flag string) {
	fields := []string{}
	cObject := countries.GetCountry(country)
	// fields = append(fields, lipgloss.NewStyle().
	// 	Background(lipgloss.Color(cObject.BgColor.Color1)).
	// 	Foreground(lipgloss.Color(cObject.FgColor.Color1)).
	// 	Render(string(" ")))

	// fields = append(fields, lipgloss.NewStyle().
	// 	Background(lipgloss.Color(cObject.BgColor.Color2)).
	// 	Foreground(lipgloss.Color(cObject.FgColor.Color2)).
	// 	Render(string(" ")))

	// fields = append(fields, lipgloss.NewStyle().
	// 	Background(lipgloss.Color(cObject.BgColor.Color3)).
	// 	Foreground(lipgloss.Color(cObject.FgColor.Color3)).
	// 	Render(string(" ")))
	fields = append(fields, cObject.Flag)
	fields = append(fields, " ", country)
	return lipgloss.JoinHorizontal(lipgloss.Left, fields...)
}

package tui

import (
	"fmt"
	"math"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lucas-ingemar/f1sched/internal/shared"
)

// type sessionState uint

const (
	// defaultTime              = time.Minute
	// timerView   sessionState = iota
	// spinnerView
	cardWidth  = 40
	cardHeight = 12
)

// var (
// 	helpStyle = lipgloss.NewStyle().Foreground(lipgloss.ANSIColor(8))
// )

type seasonScheduleModel struct {
	races          []RaceSummaryModel
	selectedIndex  int
	row            int
	col            int
	terminalWidth  int
	terminalHeight int
	cardsPerRow    int
	totalRows      int
}

func newSeasonScheduleModel(raceSchedule shared.RaceSchedule) seasonScheduleModel {
	m := seasonScheduleModel{}
	for _, r := range raceSchedule.Races {
		m.races = append(m.races, NewRaceSummary(r))
	}
	return m
}

func (m seasonScheduleModel) Init() tea.Cmd {
	models := []tea.Cmd{}
	for _, t := range m.races {
		models = append(models, t.Init())
	}
	return tea.Batch(models...)
}

func (m *seasonScheduleModel) moveFocus(direction string) error {
	switch direction {
	case "left":
		m.col -= 1
		if m.col < 0 {
			m.col = m.cardsPerRow - 1
		}

	case "right":
		m.col += 1
		if m.col > m.cardsPerRow-1 {
			m.col = 0
		}

	case "up":
		m.row -= 1
		if m.row < 0 {
			m.row = m.totalRows - 1
		}

	case "down":
		m.row += 1
		if m.row > m.totalRows-1 {
			m.row = 0
		}
	}

	m.selectedIndex = m.row*m.cardsPerRow + m.col
	if m.selectedIndex >= len(m.races) {
		m.selectedIndex = len(m.races) - 1
	}

	return nil
}

func (m seasonScheduleModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.terminalWidth = msg.Width
		m.terminalHeight = msg.Height
		m.cardsPerRow = int(m.terminalWidth / cardWidth)
		m.totalRows = int(math.Ceil(float64(len(m.races)) / float64(m.cardsPerRow)))
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "down", "left", "right":
			err := m.moveFocus(msg.String())
			_ = err
		case "ctrl+c", "q":
			return m, tea.Quit
			// case "tab":
			// 	if m.state == timerView {
			// 		m.state = spinnerView
			// 	} else {
			// 		m.state = timerView
			// 	}
			// case "n":
			// 	if m.state == timerView {
			// 		m.timer = timer.New(defaultTime)
			// 		cmds = append(cmds, m.timer.Init())
			// 	} else {
			// 		// m.Next()
			// 		// m.resetSpinner()
			// 		cmds = append(cmds, m.spinner.Tick)
			// 	}
		}
		// switch m.state {
		// // update whichever model is focused
		// case spinnerView:
		// 	m.spinner, cmd = m.spinner.Update(msg)
		// 	cmds = append(cmds, cmd)
		// default:
		// 	m.timer, cmd = m.timer.Update(msg)
		// 	cmds = append(cmds, cmd)
		// }
		// case spinner.TickMsg:
		// 	m.spinner, cmd = m.spinner.Update(msg)
		// 	cmds = append(cmds, cmd)
		// case timer.TickMsg:
		// 	m.timer, cmd = m.timer.Update(msg)
		// 	cmds = append(cmds, cmd)
	}
	return m, tea.Batch(cmds...)
}

func (m seasonScheduleModel) View() string {
	var s string

	if m.cardsPerRow == 0 {
		return ""
	}

	hList := []string{}
	vList := []string{}

	startIdx := m.row * m.cardsPerRow
	maxRows := int(math.Ceil(float64(m.terminalHeight) / float64(cardHeight+2)))
	endIdx := startIdx + maxRows*m.cardsPerRow

	trimVal := m.terminalHeight%(cardHeight+2) - 2

	row := 0
	for idx, card := range m.races {
		if idx < startIdx || idx >= endIdx {
			continue
		}
		if idx != startIdx && idx%m.cardsPerRow == 0 {
			vList = append(vList, lipgloss.JoinHorizontal(lipgloss.Top, hList...))
			hList = []string{}
			row += 1
		}
		style := getCardStyle(idx == m.selectedIndex)

		renderedCard := style.Render(fmt.Sprintf("%4s", card.View()))
		if row == maxRows-1 {
			renderedCard = strings.Join(strings.Split(renderedCard, "\n")[:trimVal], "\n")
		}
		hList = append(hList, renderedCard)

	}

	if len(hList) > 0 {
		vList = append(vList, lipgloss.JoinHorizontal(lipgloss.Top, hList...))
	}

	s += lipgloss.JoinVertical(lipgloss.Top, vList...)
	// s += helpStyle.Render("\ntab: focus next • n: new %s • q: exit\n")
	return s
}

// func Run(raceSchedule shared.RaceSchedule) error {
// 	p := tea.NewProgram(newModel(raceSchedule), tea.WithAltScreen())
// 	_, err := p.Run()
// 	return err
// }

func getCardStyle(focused bool) lipgloss.Style {
	modelStyle := lipgloss.NewStyle().
		Width(cardWidth).
		Height(cardHeight).
		Align(lipgloss.Top, lipgloss.Left).
		BorderForeground(lipgloss.ANSIColor(8)).
		BorderStyle(lipgloss.RoundedBorder())

	if focused {
		modelStyle = modelStyle.
			BorderForeground(lipgloss.ANSIColor(5))
	}
	return modelStyle
}

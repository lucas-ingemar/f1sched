package main

import (
	"fmt"
	"log"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lucas-ingemar/f1sched/internal/bubbletea/bubbles"
	"github.com/lucas-ingemar/f1sched/internal/ergast"
	"github.com/lucas-ingemar/f1sched/internal/shared"
)

/*
This example assumes an existing understanding of commands and messages. If you
haven't already read our tutorials on the basics of Bubble Tea and working
with commands, we recommend reading those first.

Find them at:
https://github.com/charmbracelet/bubbletea/tree/master/tutorials/commands
https://github.com/charmbracelet/bubbletea/tree/master/tutorials/basics
*/

// sessionState is used to track which model is focused
type sessionState uint

const (
	defaultTime              = time.Minute
	timerView   sessionState = iota
	spinnerView
)

var (
	// Available spinners
	spinners = []spinner.Spinner{
		spinner.Line,
		spinner.Dot,
		spinner.MiniDot,
		spinner.Jump,
		spinner.Pulse,
		spinner.Points,
		spinner.Globe,
		spinner.Moon,
		spinner.Monkey,
	}
	modelStyle = lipgloss.NewStyle().
			Width(30).
			Height(11).
			Align(lipgloss.Top, lipgloss.Left).
			BorderStyle(lipgloss.NormalBorder())
	focusedModelStyle = lipgloss.NewStyle().
				Width(15).
				Height(5).
				Align(lipgloss.Center, lipgloss.Center).
				BorderStyle(lipgloss.NormalBorder()).
				BorderForeground(lipgloss.Color("69"))
	spinnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("69"))
	helpStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
)

type mainModel struct {
	state   sessionState
	timer   timer.Model
	timers  []bubbles.RaceSummaryModel
	spinner spinner.Model
	index   int
}

func newModel(timeout time.Duration, races []shared.Race) mainModel {
	m := mainModel{state: timerView}
	m.timer = timer.New(timeout)
	m.spinner = spinner.New()

	for _, r := range races {
		m.timers = append(m.timers, bubbles.NewRaceSummary(r))
	}
	return m
}

func (m mainModel) Init() tea.Cmd {
	// start the timer and spinner on program start
	models := []tea.Cmd{
		m.timer.Init(), m.spinner.Tick,
	}
	for _, t := range m.timers {
		models = append(models, t.Init())
	}
	return tea.Batch(models...)
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		fmt.Println(msg.Width, msg.Height)
	case tea.KeyMsg:
		// fmt.Println(msg.String())
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab":
			if m.state == timerView {
				m.state = spinnerView
			} else {
				m.state = timerView
			}
		case "n":
			if m.state == timerView {
				m.timer = timer.New(defaultTime)
				cmds = append(cmds, m.timer.Init())
			} else {
				m.Next()
				m.resetSpinner()
				cmds = append(cmds, m.spinner.Tick)
			}
		}
		switch m.state {
		// update whichever model is focused
		case spinnerView:
			m.spinner, cmd = m.spinner.Update(msg)
			cmds = append(cmds, cmd)
		default:
			m.timer, cmd = m.timer.Update(msg)
			cmds = append(cmds, cmd)
		}
	case spinner.TickMsg:
		m.spinner, cmd = m.spinner.Update(msg)
		cmds = append(cmds, cmd)
	case timer.TickMsg:
		m.timer, cmd = m.timer.Update(msg)
		cmds = append(cmds, cmd)
	}
	return m, tea.Batch(cmds...)
}

func (m mainModel) View() string {
	var s string
	model := m.currentFocusedModel()
	if m.state == timerView {
		s += lipgloss.JoinHorizontal(lipgloss.Top, focusedModelStyle.Render(fmt.Sprintf("%4s", m.timer.View())), modelStyle.Render(m.spinner.View()))
	} else {

		v := []string{
			modelStyle.Render(fmt.Sprintf("%4s", m.timer.View())), focusedModelStyle.Render(m.spinner.View()),
		}
		for _, t := range m.timers {
			v = append(v, modelStyle.Render(fmt.Sprintf("%4s", t.View())))
		}
		s += lipgloss.JoinHorizontal(lipgloss.Top, v...)
	}
	s += helpStyle.Render(fmt.Sprintf("\ntab: focus next • n: new %s • q: exit\n", model))
	return s
}

func (m mainModel) currentFocusedModel() string {
	if m.state == timerView {
		return "timer"
	}
	return "spinner"
}

func (m *mainModel) Next() {
	if m.index == len(spinners)-1 {
		m.index = 0
	} else {
		m.index++
	}
}

func (m *mainModel) resetSpinner() {
	m.spinner = spinner.New()
	m.spinner.Style = spinnerStyle
	m.spinner.Spinner = spinners[m.index]
}

func main() {
	races, err := ergast.GetRaceData()
	if err != nil {
		panic(err)
	}

	p := tea.NewProgram(newModel(defaultTime, races))
	// For fullscreen
	// p := tea.NewProgram(newModel(defaultTime), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

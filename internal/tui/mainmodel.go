package tui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lucas-ingemar/f1sched/internal/shared"
	"github.com/muesli/ansi"
)

// type sessionState uint

const (
// defaultTime              = time.Minute
// timerView   sessionState = iota
// spinnerView
// cardWidth  = 40
// cardHeight = 12
)

var (
	inactiveTabBorder = tabBorderWithBottom("┴", "─", "┴")
	activeTabBorder   = tabBorderWithBottom("┘", " ", "└")
	paddingBorder     = generatePaddingBorder()
	docStyle          = lipgloss.NewStyle().Padding(0, 0, 0, 0)
	highlightColor    = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	inactiveTabStyle  = lipgloss.NewStyle().Border(inactiveTabBorder, true).BorderForeground(highlightColor).Padding(0, 1)
	activeTabStyle    = inactiveTabStyle.Copy().Border(activeTabBorder, true)
	windowStyle       = lipgloss.NewStyle().BorderForeground(highlightColor).Padding(2, 0, 0, 0).Align(lipgloss.Center).Border(lipgloss.NormalBorder()).UnsetBorderTop()
	helpStyle         = lipgloss.NewStyle().Foreground(lipgloss.ANSIColor(8))
)

type mainModel struct {
	Tabs       []string
	TabContent []tea.Model
	// ssm            tea.Model
	activeTab      int
	terminalWidth  int
	terminalHeight int
}

func newModel(raceSchedule shared.RaceSchedule, driverStandings shared.DriverStandings) mainModel {
	// m := mainModel{}
	// seasonScheduleModel := newSeasonScheduleModel(raceSchedule)
	tabs := []string{"Season Schedule", "Driver Standing", "Team Standing"}
	// tabContent := []string{"Lip Gloss Tab", "Mascara Tab", "Foundation Tab"}
	tabContent := []tea.Model{
		newSeasonScheduleModel(raceSchedule),
		newDriverStandingsModel(driverStandings),
		newDriverStandingsModel(driverStandings),
	}
	m := mainModel{Tabs: tabs, TabContent: tabContent}
	return m
}

func (m mainModel) Init() tea.Cmd {
	models := []tea.Cmd{}
	// models = append(models, m.ssm.Init())
	for _, t := range m.TabContent {
		models = append(models, t.Init())
	}
	return tea.Batch(models...)
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.terminalWidth = msg.Width
		m.terminalHeight = msg.Height
		msg.Width = msg.Width - windowStyle.GetHorizontalFrameSize() - 6
		msg.Height = msg.Height - windowStyle.GetVerticalFrameSize() - 1
		for idx, tm := range m.TabContent {
			m.TabContent[idx], cmd = tm.Update(msg)
			cmds = append(cmds, cmd)
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab":
			if m.activeTab == len(m.Tabs)-1 {
				m.activeTab = 0
			} else {
				m.activeTab = min(m.activeTab+1, len(m.Tabs)-1)
			}
			return m, nil
		case "shift+tab":
			m.activeTab = max(m.activeTab-1, 0)
			return m, nil
			// if m.state == timerView {
			// 	m.state = spinnerView
			// } else {
			// 	m.state = timerView
			// }
		default:
			m.TabContent[m.activeTab], cmd = m.TabContent[m.activeTab].Update(msg)
			// m.ssm, cmd = m.ssm.Update(msg)
			cmds = append(cmds, cmd)
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
	// m.ssm, cmd = m.ssm.Update(msg)
	// cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m mainModel) View() string {
	doc := strings.Builder{}

	var renderedTabs []string

	totalWidth := 0
	for i, t := range m.Tabs {
		var style lipgloss.Style
		isFirst, isLast, isActive := i == 0, i == len(m.Tabs)-1, i == m.activeTab
		_ = isLast
		if isActive {
			style = activeTabStyle.Copy()
		} else {
			style = inactiveTabStyle.Copy()
		}
		border, _, _, _, _ := style.GetBorder()
		if isFirst && isActive {
			border.BottomLeft = "│"
		} else if isFirst && !isActive {
			border.BottomLeft = "├"
		}
		style = style.Border(border)
		renderedTab := style.Render(t)

		totalWidth += ansi.PrintableRuneWidth(strings.Split(renderedTab, "\n")[0])

		renderedTabs = append(renderedTabs, renderedTab)
	}

	paddingWidth := m.terminalWidth - totalWidth - 4
	if paddingWidth > 0 {
		style := inactiveTabStyle.Copy().Border(paddingBorder, true, true, true, true)
		paddingTab := style.Render(strings.Repeat(" ", paddingWidth-11) + "2024 Season")
		renderedTabs = append(renderedTabs, paddingTab)
	}

	row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)
	doc.WriteString(row)
	// if m.Tabs[m.activeTab] == "Season Schedule" {
	doc.WriteString(
		windowStyle.Width((m.terminalWidth - windowStyle.GetHorizontalFrameSize())).
			Height(m.terminalHeight - windowStyle.GetVerticalFrameSize() - 1).
			Render(m.TabContent[m.activeTab].View()))
	// } else {
	// 	doc.WriteString(
	// 		windowStyle.Width((m.terminalWidth - windowStyle.GetHorizontalFrameSize())).
	// 			Height(m.terminalHeight - windowStyle.GetVerticalFrameSize() - 1).
	// 			Render(m.TabContent[m.activeTab]))
	// }

	doc.WriteString(helpStyle.Render("\ntab: focus next • n: new %s • q: exit"))
	return docStyle.Render(doc.String())

	// var s string

	// vList := []string{}

	// s += lipgloss.JoinVertical(lipgloss.Top, vList...)
	// s += helpStyle.Render("\ntab: focus next • n: new %s • q: exit\n")
	// return s
}

func tabBorderWithBottom(left, middle, right string) lipgloss.Border {
	border := lipgloss.RoundedBorder()
	border.BottomLeft = left
	border.Bottom = middle
	border.BottomRight = right
	return border
}

func generatePaddingBorder() lipgloss.Border {
	// border, _, _, _, _ := style.GetBorder()
	border := lipgloss.NormalBorder()
	border.Top = ""
	border.Left = ""
	border.Right = ""
	border.TopLeft = ""
	border.TopRight = ""
	border.BottomLeft = border.Bottom
	border.BottomRight = "┐"
	return border
}

func Run(raceSchedule shared.RaceSchedule, driverStandings shared.DriverStandings) error {
	p := tea.NewProgram(newModel(raceSchedule, driverStandings), tea.WithAltScreen())
	_, err := p.Run()
	return err
}

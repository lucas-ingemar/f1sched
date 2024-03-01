package bubbles

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lucas-ingemar/f1sched/internal/shared"
)

type RaceSummaryModel struct {
	// Spinner settings to use. See type Spinner.
	// Spinner Spinner

	// Style sets the styling for the spinner. Most of the time you'll just
	// want foreground and background coloring, and potentially some padding.
	//
	// For an introduction to styling with Lip Gloss see:
	// https://github.com/charmbracelet/lipgloss
	Style lipgloss.Style

	race shared.Race

	frame int
	id    int
	tag   int
}

func (m RaceSummaryModel) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}
func (m RaceSummaryModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// switch msg := msg.(type) {

	// // Is it a key press?
	// case tea.KeyMsg:

	//     // Cool, what was the actual key pressed?
	//     switch msg.String() {

	//     // These keys should exit the program.
	//     case "ctrl+c", "q":
	//         return m, tea.Quit

	//     // The "up" and "k" keys move the cursor up
	//     case "up", "k":
	//         if m.cursor > 0 {
	//             m.cursor--
	//         }

	//     // The "down" and "j" keys move the cursor down
	//     case "down", "j":
	//         if m.cursor < len(m.choices)-1 {
	//             m.cursor++
	//         }

	//     // The "enter" key and the spacebar (a literal space) toggle
	//     // the selected state for the item that the cursor is pointing at.
	//     case "enter", " ":
	//         _, ok := m.selected[m.cursor]
	//         if ok {
	//             delete(m.selected, m.cursor)
	//         } else {
	//             m.selected[m.cursor] = struct{}{}
	//         }
	//     }
	// }

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m RaceSummaryModel) View() string {
	// The header
	s := ""
	s += lipgloss.NewStyle().
		Background(lipgloss.Color("#CE1126")).
		Foreground(lipgloss.Color("#FFFFFF")).
		// Foreground(lipgloss.Color("#CE1126")).
		// Background(lipgloss.Color("#FFFFFF")).
		Align(lipgloss.Center).
		Width(24).
		Render(m.race.Country)
	s += "\n      "
	s += lipgloss.NewStyle().
		// Background(lipgloss.Color("#CE1126")).
		// Foreground(lipgloss.Color("#FFFFFF")).
		Foreground(lipgloss.Color("#CE1126")).
		Background(lipgloss.Color("#FFFFFF")).
		Align(lipgloss.Center).
		Width(24).
		Render(m.race.Circuit)
	// s += "\n"
	s += "\n   "
	s += lipgloss.NewStyle().
		Background(lipgloss.Color("#CE1126")).
		Foreground(lipgloss.Color("#FFFFFF")).
		// Foreground(lipgloss.Color("#CE1126")).
		// Background(lipgloss.Color("#FFFFFF")).
		Align(lipgloss.Center).
		Width(24).
		Render(m.race.RaceName)
	s += "\n"
	s += "\n"
	s += "Feb 29 - Mar 02, 2023"
	s += "\n"
	s += "\n"

	// s += lipgloss.NewStyle().
	// 	Align(lipgloss.Right).
	// 	Width(30).
	// 	Render("Practice 1: 29 Feb, 15:00")
	s += "Practice 1:      29 Feb, 15:00"
	s += "\n"
	s += "Practice 2:      29 Feb, 21:00"
	s += "\n"
	s += "Practice 3:      01 Mar, 15:00"
	s += "\n"
	s += "\n"
	s += "Qualifying:      01 Mar, 21:00"
	s += "\n"
	s += "Race:            02 Mar, 15:00"

	// // Iterate over our choices
	// for i, choice := range m.choices {

	//     // Is the cursor pointing at this choice?
	//     cursor := " " // no cursor
	//     if m.cursor == i {
	//         cursor = ">" // cursor!
	//     }

	//     // Is this choice selected?
	//     checked := " " // not selected
	//     if _, ok := m.selected[i]; ok {
	//         checked = "x" // selected!
	//     }

	//     // Render the row
	//     s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	// }

	// The footer
	// s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}

// New returns a model with default values.
func NewRaceSummary(race shared.Race) RaceSummaryModel {
	m := RaceSummaryModel{
		race: race,
		// Spinner: Line,
		// id:      nextID(),
	}

	// for _, opt := range opts {
	// 	opt(&m)
	// }

	return m
}

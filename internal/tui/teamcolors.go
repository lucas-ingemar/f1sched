package tui

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	tcRedBullRacing = lipgloss.NewStyle().Background(lipgloss.Color("#336ec0")).Foreground(lipgloss.Color("#FFFFFF"))
	tcFerrari       = lipgloss.NewStyle().Background(lipgloss.Color("#e60023")).Foreground(lipgloss.Color("#FFFFFF"))
	tcMercedes      = lipgloss.NewStyle().Background(lipgloss.Color("#27f1cf")).Foreground(lipgloss.Color("#000000"))
	tcMcLaren       = lipgloss.NewStyle().Background(lipgloss.Color("#fd7d00")).Foreground(lipgloss.Color("#FFFFFF"))
	tcAstonMartin   = lipgloss.NewStyle().Background(lipgloss.Color("#209671")).Foreground(lipgloss.Color("#FFFFFF"))
	tcSauber        = lipgloss.NewStyle().Background(lipgloss.Color("#51df50")).Foreground(lipgloss.Color("#000000"))
	tcHaas          = lipgloss.NewStyle().Background(lipgloss.Color("#b4b7ba")).Foreground(lipgloss.Color("#000000"))
	tcRb            = lipgloss.NewStyle().Background(lipgloss.Color("#6490fa")).Foreground(lipgloss.Color("#000000"))
	tcWilliams      = lipgloss.NewStyle().Background(lipgloss.Color("#62c1fd")).Foreground(lipgloss.Color("#000000"))
	tcAlpine        = lipgloss.NewStyle().Background(lipgloss.Color("#fc85b8")).Foreground(lipgloss.Color("#000000"))
)

func getTeamStyle(teamName string) lipgloss.Style {
	switch teamName {
	case "Red Bull Racing Honda RBPT":
		return tcRedBullRacing
	case "Ferrari":
		return tcFerrari
	case "Mercedes":
		return tcMercedes
	case "McLaren Mercedes":
		return tcMcLaren
	case "Aston Martin Aramco Mercedes":
		return tcAstonMartin
	case "Kick Sauber Ferrari":
		return tcSauber
	case "Haas Ferrari":
		return tcHaas
	case "RB Honda RBPT":
		return tcRb
	case "Williams Mercedes":
		return tcWilliams
	case "Alpine Renault":
		return tcAlpine
	default:
		return lipgloss.NewStyle()
	}
}

package tui

import "github.com/charmbracelet/lipgloss"

var (
	SubtleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))

	// feedback
	CorrectStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#32CD32")) // Verde
	IncorrectStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF4136")) // Rojo

	// borde de la ventana
	NormalBorder    = lipgloss.NewStyle().Border(lipgloss.NormalBorder(), true).BorderForeground(lipgloss.Color("240"))
	CorrectBorder   = lipgloss.NewStyle().Border(lipgloss.NormalBorder(), true).BorderForeground(CorrectStyle.GetForeground())
	IncorrectBorder = lipgloss.NewStyle().Border(lipgloss.NormalBorder(), true).BorderForeground(IncorrectStyle.GetForeground())
)

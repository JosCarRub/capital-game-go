package components

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Width    int
	Progress float64
	Style    lipgloss.Style
}

func NewProgressBar(width int) Model {
	return Model{
		Width: width,

		Style: lipgloss.NewStyle().Foreground(lipgloss.Color("#80e895")),
	}
}

func (m *Model) SetProgress(progress float64) {
	if progress < 0.0 {
		m.Progress = 0.0
	} else if progress > 1.0 {
		m.Progress = 1.0
	} else {
		m.Progress = progress
	}
}

func (m Model) View() string {
	filledWidth := int(m.Progress * float64(m.Width))
	bar := strings.Repeat("█", filledWidth) + strings.Repeat("─", m.Width-filledWidth)

	return m.Style.Render(fmt.Sprintf("[%s]", bar))
}

package views

import (
	"capital-game-go/internal/tui/components"
	"capital-game-go/internal/tui/style"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MainMenuModel struct {
	cursor  int
	choices []string
	width   int
	height  int
}

func (m *MainMenuModel) SetSize(width, height int) {
	m.width = width
	m.height = height
}

func NewMainMenu() MainMenuModel {
	return MainMenuModel{
		choices: []string{"Jugar", "Ranking", "Salir"},
	}
}

func (m MainMenuModel) Init() tea.Cmd {
	return nil
}

func (m MainMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, key.NewBinding(key.WithKeys("up", "k"))):
			if m.cursor > 0 {
				m.cursor--
			}
		case key.Matches(msg, key.NewBinding(key.WithKeys("down", "j"))):
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case key.Matches(msg, key.NewBinding(key.WithKeys("enter"))):
			selectedItem := m.choices[m.cursor]
			switch selectedItem {
			case "Jugar":
				return m, func() tea.Msg {
					return SwitchToViewMsg{NewView: GameView}
				}
			case "Ranking":
				return m, func() tea.Msg {
					return SwitchToViewMsg{NewView: LeaderboardView}
				}
			case "Salir":
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

func (m MainMenuModel) View() string {
	var b strings.Builder

	b.WriteString(components.View())
	b.WriteString("\n\n")

	for i, choice := range m.choices {
		cursor := "  "
		style := lipgloss.NewStyle()
		if m.cursor == i {
			cursor = "▶ "
			style = style.Foreground(lipgloss.Color("#00AEEF")).Bold(true)
		}
		b.WriteString(style.Render(cursor + choice))
		b.WriteString("\n")
	}

	b.WriteString("\n⭐ Un juego de JosCarRub ⭐\n")

	helpView := style.HelpStyle.Render("↑/↓: Moverse  •  Enter: Seleccionar  •  Ctrl+C: Salir")

	mainContent := lipgloss.JoinVertical(lipgloss.Center, b.String(), helpView)

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, mainContent)
}

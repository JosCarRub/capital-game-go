package views

import (
	"capital-game-go/internal/tui/components"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type MainMenuModel struct {
	cursor  int
	choices []string
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
	s := components.View() + "\n\n"

	for i, choice := range m.choices {
		cursor := "  "
		if m.cursor == i {
			cursor = "▶ "
		}
		s += cursor + choice + "\n"
	}

	s += "\n⭐ Un juego de JosCarRub ⭐\n"
	return s
}

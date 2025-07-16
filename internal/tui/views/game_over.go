package views

import (
	"capital-game-go/internal/tui/style"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ScoreSubmittedMsg struct {
	PlayerName string
	Hits       int
}

type GameOverModel struct {
	hits           int
	misses         int
	totalQuestions int
	textInput      textinput.Model
	isSubmitting   bool
	width          int
	height         int
}

func (m *GameOverModel) SetSize(width, height int) {
	m.width = width
	m.height = height
}

func NewGameOverView(hits, misses, total int) GameOverModel {
	ti := textinput.New()
	ti.Placeholder = "Introduce tu nombre"
	ti.Focus()
	ti.CharLimit = 20
	ti.Width = 20

	return GameOverModel{
		hits:           hits,
		misses:         misses,
		totalQuestions: total,
		textInput:      ti,
	}
}

func (m GameOverModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m GameOverModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	if m.isSubmitting {
		return m, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, key.NewBinding(key.WithKeys("enter"))) {
			m.isSubmitting = true
			return m, func() tea.Msg {
				return ScoreSubmittedMsg{PlayerName: m.textInput.Value(), Hits: m.hits}
			}
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m GameOverModel) View() string {
	var b strings.Builder

	title := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FFDC00")).Render("¡Partida Finalizada!")
	b.WriteString(title)
	b.WriteString("\n\n")

	scoreText := fmt.Sprintf("Tu puntuación: %d / %d", m.hits, m.totalQuestions)
	scoreStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#32CD32"))
	b.WriteString(scoreStyle.Render(scoreText))
	b.WriteString("\n\n")

	b.WriteString("Introduce tu nombre para guardar la puntuación:\n")
	b.WriteString(m.textInput.View())
	b.WriteString("\n\n")

	promoBoxStyle := lipgloss.NewStyle().
		Border(lipgloss.DoubleBorder(), true).
		BorderForeground(lipgloss.Color("#555")).
		Padding(1, 2)

	promoText := "Si te ha gustado el juego, ¡apoya el proyecto!\n"
	promoLink := lipgloss.NewStyle().Foreground(lipgloss.Color("#00AEEF")).Render("⭐ Dame una estrella en GitHub ⭐")
	b.WriteString(promoBoxStyle.Render(promoText + promoLink))
	b.WriteString("\n\n")

	helpView := style.HelpStyle.Render("Enter: Guardar y ver ranking  •  Esc: Volver al menú")
	b.WriteString(helpView)

	content := lipgloss.JoinVertical(lipgloss.Center, b.String())
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, content)
}

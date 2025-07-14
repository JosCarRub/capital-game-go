package views

import (
	"capital-game-go/internal/game"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type GameOverMsg struct {
	Hits   int
	Misses int
}

type GameViewModel struct {
	gameSession     *game.Game
	currentQuestion *game.Country
	textInput       textinput.Model
	isGameOver      bool
	RoundSize       int
}

func NewGameView(countries []game.Country) GameViewModel {
	ti := textinput.New()
	ti.Placeholder = "Escribe la capital aquí..."
	ti.Focus()
	ti.CharLimit = 50
	ti.Width = 30

	session := game.NewGame(countries)
	question, _ := session.NextQuestion()

	return GameViewModel{
		gameSession:     session,
		currentQuestion: question,
		textInput:       ti,
		isGameOver:      false,
		RoundSize:       session.RoundSize,
	}
}

func (m GameViewModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m GameViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	if m.isGameOver {
		return m, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			userInput := m.textInput.Value()
			if game.CheckAnswer(userInput, m.currentQuestion.Capital) {
				m.gameSession.RecordHit()
			} else {
				m.gameSession.RecordMiss()
			}

			nextQ, hasNext := m.gameSession.NextQuestion()
			if !hasNext || m.gameSession.Hits+m.gameSession.Misses >= m.gameSession.RoundSize {
				m.isGameOver = true
				return m, func() tea.Msg {

					return GameOverMsg{Hits: m.gameSession.Hits, Misses: m.gameSession.Misses}
				}
			}

			m.currentQuestion = nextQ
			m.textInput.Reset()
			return m, nil
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m GameViewModel) View() string {
	if m.isGameOver {
		return "Calculando resultados finales..."
	}

	var b strings.Builder

	header := fmt.Sprintf("Pregunta %d/%d", m.gameSession.Hits+m.gameSession.Misses+1, m.gameSession.RoundSize)
	b.WriteString(header + "\n\n")

	b.WriteString("¿Cuál es la capital de...\n")
	b.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("#00AEEF")).Bold(true).Render(m.currentQuestion.Name))
	b.WriteString("\n\n")

	b.WriteString(m.textInput.View())
	b.WriteString("\n\n")

	hitsText := fmt.Sprintf("Aciertos: %d ✅", m.gameSession.Hits)
	missesText := fmt.Sprintf("Fallos: %d ❌", m.gameSession.Misses)

	progress := float64(m.gameSession.Hits+m.gameSession.Misses) / float64(m.gameSession.RoundSize)

	if progress > 1.0 {
		progress = 1.0
	}
	progressWidth := 20
	filledWidth := int(progress * float64(progressWidth))
	progressText := fmt.Sprintf("[%s%s]", strings.Repeat("█", filledWidth), strings.Repeat("─", progressWidth-filledWidth))

	statusBar := lipgloss.JoinHorizontal(lipgloss.Top,
		lipgloss.NewStyle().Width(20).Render(hitsText),
		lipgloss.NewStyle().Width(25).Align(lipgloss.Center).Render("Progreso "+progressText),
		lipgloss.NewStyle().Width(20).Align(lipgloss.Right).Render(missesText),
	)
	b.WriteString(statusBar)

	return b.String()
}

package views

import (
	"capital-game-go/internal/database"
	"capital-game-go/internal/game"
	"capital-game-go/internal/tui/style"
	"database/sql"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type leaderboardLoadedMsg struct{ scores []game.PlayerScore }
type leaderboardErrorMsg struct{ err error }

type LeaderboardModel struct {
	db        *sql.DB
	spinner   spinner.Model
	scores    []game.PlayerScore
	isLoading bool
	err       error
	width     int
	height    int
}

func (m *LeaderboardModel) SetSize(width, height int) {
	m.width = width
	m.height = height
}

func NewLeaderboardView(db *sql.DB) LeaderboardModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#00AEEF"))
	return LeaderboardModel{
		db:        db,
		spinner:   s,
		isLoading: true,
	}
}

func (m LeaderboardModel) loadLeaderboard() tea.Msg {
	scores, err := database.GetLeaderboard(m.db)
	if err != nil {
		return leaderboardErrorMsg{err: err}
	}
	return leaderboardLoadedMsg{scores: scores}
}

func (m LeaderboardModel) Init() tea.Cmd {
	return tea.Batch(m.loadLeaderboard, m.spinner.Tick)
}

func (m LeaderboardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case leaderboardLoadedMsg:
		m.isLoading = false
		m.scores = msg.scores
		m.err = nil
		return m, nil
	case leaderboardErrorMsg:
		m.isLoading = false
		m.err = msg.err
		return m, nil
	}

	var cmd tea.Cmd
	if m.isLoading {
		m.spinner, cmd = m.spinner.Update(msg)
	}
	return m, cmd
}

func (m LeaderboardModel) View() string {
	var content string

	if m.isLoading {
		content = fmt.Sprintf("\n   %s Cargando ranking...\n\n", m.spinner.View())
	} else if m.err != nil {
		content = fmt.Sprintf("\nError al cargar el ranking:\n%v\n\nPulsa 'Esc' para volver.", m.err)
	} else {
		var b strings.Builder

		title := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FFDC00")).Render("🏆 RANKING DE JUGADORES 🏆")
		b.WriteString(title)
		b.WriteString("\n\n")

		headerStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#00AEEF"))
		header := lipgloss.JoinHorizontal(lipgloss.Top,
			headerStyle.Copy().Width(8).Render("#"),
			headerStyle.Copy().Width(30).Render("Jugador"),
			headerStyle.Copy().Width(15).Align(lipgloss.Right).Render("Puntos"),
		)
		b.WriteString(header)
		b.WriteString("\n")
		b.WriteString(strings.Repeat("─", 53))
		b.WriteString("\n")

		for i, score := range m.scores {
			if i >= 10 {
				break
			}
			rank := fmt.Sprintf("%d", i+1)
			rankStyle := lipgloss.NewStyle().Width(8)
			if i == 0 {
				rank = "🥇"
			} else if i == 1 {
				rank = "🥈"
			} else if i == 2 {
				rank = "🥉"
			}

			row := lipgloss.JoinHorizontal(lipgloss.Top,
				rankStyle.Render(rank),
				lipgloss.NewStyle().Width(30).Render(score.Name),
				lipgloss.NewStyle().Width(15).Align(lipgloss.Right).Render(fmt.Sprintf("%d", score.Points)),
			)
			b.WriteString(row)
			b.WriteString("\n")
		}

		b.WriteString("\n\n")
		instructions := style.HelpStyle.Render("Pulsa 'Esc' para volver al menú principal.")
		b.WriteString(instructions)

		content = b.String()
	}

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, content)
}

package tui

import (
	"capital-game-go/internal/database"
	"capital-game-go/internal/game"
	"capital-game-go/internal/tui/views"
	"database/sql"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var fadeRamp = []string{"#FFFFFF", "#E0E0E0", "#C0C0C0", "#A0A0A0", "#808080", "#606060", "#404040", "#202020", "#000000"}

const fadeDuration = 50 * time.Millisecond

type fadeOutMsg struct{}

type MainModel struct {
	db          *sql.DB
	view        views.CurrentView
	countries   []game.Country
	menu        views.MainMenuModel
	game        views.GameViewModel
	gameOver    views.GameOverModel
	leaderboard views.LeaderboardModel
	width       int
	height      int
	err         error

	isFadingOut bool
	fadeStep    int
	nextView    views.CurrentView
}

func NewMainModel(db *sql.DB, countries []game.Country) MainModel {
	return MainModel{
		db:          db,
		view:        views.MainMenu,
		countries:   countries,
		menu:        views.NewMainMenu(),
		game:        views.NewGameView(countries),
		leaderboard: views.NewLeaderboardView(db),
	}
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m *MainModel) startFade(targetView views.CurrentView) (tea.Model, tea.Cmd) {
	m.isFadingOut = true
	m.fadeStep = 0
	m.nextView = targetView
	return m, func() tea.Msg {
		time.Sleep(fadeDuration)
		return fadeOutMsg{}
	}
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.isFadingOut {
		switch msg.(type) {
		case fadeOutMsg:
			m.fadeStep++
			if m.fadeStep >= len(fadeRamp) {
				m.isFadingOut = false
				m.view = m.nextView
				if m.view == views.LeaderboardView {
					m.leaderboard = views.NewLeaderboardView(m.db)
					m.leaderboard.SetSize(m.width, m.height)
					return m, m.leaderboard.Init()
				}
				return m, nil
			}
			return m, func() tea.Msg {
				time.Sleep(fadeDuration)
				return fadeOutMsg{}
			}
		case tea.KeyMsg:

			return m, nil
		}
	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.menu.SetSize(m.width, m.height)
		m.game.SetSize(m.width, m.height)
		m.gameOver.SetSize(m.width, m.height)
		m.leaderboard.SetSize(m.width, m.height)

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "esc":
			if m.view != views.MainMenu {
				m.game = views.NewGameView(m.countries)
				m.game.SetSize(m.width, m.height)
				return m.startFade(views.MainMenu)
			}
		}
	case views.SwitchToViewMsg:
		if msg.NewView == views.MainMenu {
			m.game = views.NewGameView(m.countries)
			m.game.SetSize(m.width, m.height)
		}
		return m.startFade(msg.NewView)

	case views.GameOverMsg:
		m.gameOver = views.NewGameOverView(msg.Hits, msg.Misses, m.game.RoundSize)
		m.gameOver.SetSize(m.width, m.height)
		return m.startFade(views.GameOverView)

	case views.ScoreSubmittedMsg:
		database.SaveScore(m.db, msg.PlayerName, msg.Hits)
		return m.startFade(views.LeaderboardView)
	}

	var cmd tea.Cmd
	switch m.view {
	case views.MainMenu:
		newModel, newCmd := m.menu.Update(msg)
		m.menu = newModel.(views.MainMenuModel)
		cmd = newCmd
	case views.GameView:
		newModel, newCmd := m.game.Update(msg)
		m.game = newModel.(views.GameViewModel)
		cmd = newCmd
	case views.GameOverView:
		newModel, newCmd := m.gameOver.Update(msg)
		m.gameOver = newModel.(views.GameOverModel)
		cmd = newCmd
	case views.LeaderboardView:
		newModel, newCmd := m.leaderboard.Update(msg)
		m.leaderboard = newModel.(views.LeaderboardModel)
		cmd = newCmd
	}

	return m, cmd
}

func (m MainModel) View() string {
	if m.width == 0 || m.height == 0 {
		return "Inicializando..."
	}

	var currentView string
	switch m.view {
	case views.MainMenu:
		currentView = m.menu.View()
	case views.GameView:
		currentView = m.game.View()
	case views.GameOverView:
		currentView = m.gameOver.View()
	case views.LeaderboardView:
		currentView = m.leaderboard.View()
	default:
		currentView = "Vista desconocida."
	}

	if m.isFadingOut {
		return lipgloss.NewStyle().
			Foreground(lipgloss.Color(fadeRamp[m.fadeStep])).
			Render(currentView)
	}

	return currentView
}

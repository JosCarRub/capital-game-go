package tui

import (
	"capital-game-go/internal/database"
	"capital-game-go/internal/game"
	"capital-game-go/internal/tui/views"
	"database/sql"

	tea "github.com/charmbracelet/bubbletea"
)

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

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
				m.view = views.MainMenu
			}
		}
	case views.SwitchToViewMsg:
		if msg.NewView == views.MainMenu {
			m.game = views.NewGameView(m.countries)
			m.game.SetSize(m.width, m.height)
		}
		m.view = msg.NewView
		if m.view == views.LeaderboardView {
			m.leaderboard = views.NewLeaderboardView(m.db)
			m.leaderboard.SetSize(m.width, m.height)
			return m, m.leaderboard.Init()
		}
		return m, nil
	case views.GameOverMsg:
		m.gameOver = views.NewGameOverView(msg.Hits, msg.Misses, m.game.RoundSize)
		m.gameOver.SetSize(m.width, m.height)
		m.view = views.GameOverView
		return m, m.gameOver.Init()
	case views.ScoreSubmittedMsg:
		database.SaveScore(m.db, msg.PlayerName, msg.Hits)
		m.view = views.LeaderboardView
		m.leaderboard = views.NewLeaderboardView(m.db)
		m.leaderboard.SetSize(m.width, m.height)
		return m, m.leaderboard.Init()
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

	switch m.view {
	case views.MainMenu:
		return m.menu.View()
	case views.GameView:
		return m.game.View()
	case views.GameOverView:
		return m.gameOver.View()
	case views.LeaderboardView:
		return m.leaderboard.View()
	default:
		return "Vista desconocida."
	}
}

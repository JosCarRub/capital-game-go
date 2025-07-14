package views

type CurrentView int

const (
	MainMenu CurrentView = iota
	GameView
	GameOverView
	LeaderboardView
)

type SwitchToViewMsg struct{ NewView CurrentView }

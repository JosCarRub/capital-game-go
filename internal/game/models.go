package game

type Country struct {
	Name    string `json:"name"`
	Capital string `json:"capital"`
}

type PlayerScore struct {
	ID     int
	Name   string
	Points int
}

package main

import (
	"capital-game-go/internal/database"
	"capital-game-go/internal/game"
	"capital-game-go/internal/tui"
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {

	db, err := database.NewConnection()
	if err != nil {
		log.Fatalf("Error fatal al conectar con la base de datos: %v", err)
	}
	defer db.Close()

	err = database.InitializeSchema(db)
	if err != nil {
		log.Fatalf("Error al inicializar el esquema de la BD: %v", err)
	}

	countries, err := game.LoadCountries("data/countries.json")
	if err != nil {
		log.Fatalf("Error al cargar los datos del juego: %v", err)
	}

	m := tui.NewMainModel(db, countries)

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Printf("¡Oh no! Hubo un error: %v", err)
		os.Exit(1)
	}
}

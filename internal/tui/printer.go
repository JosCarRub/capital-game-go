package tui

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/fatih/color"

	"capital-game-go/internal/game"
)

func PlayIntroAnimation() {
	ClearScreen()
	lines := strings.Split(SkyArt, "\n")
	for _, line := range lines {
		for _, char := range line {
			if char == '*' {

				if rand.Float64() < 0.08 {
					color.Yellow("%c", char)
				} else if rand.Float64() < 0.05 {
					color.Cyan("%c", char)
				} else {
					fmt.Printf("%c", char)
				}
			} else {
				fmt.Printf("%c", char)
			}
		}
		fmt.Println()
		time.Sleep(300 * time.Millisecond)
	}
}

// limpia pantalla
func ClearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// título e instrucciones
func ShowWelcomeScreen() {
	ClearScreen()
	color.Yellow(TitleArt)
	fmt.Println()
	color.Cyan("JUEGO CAPITALES EN GO REALIZADO POR JOSÉ CARLOS   --GITHUB: JosCarRub ---")
	fmt.Println("\nEl juego consiste en introducir la capital del país que te aparezca en la pantalla.")
	fmt.Println("\nTu puntuación quedará guardada en una Base de Datos y podrás consultar tu posición en la tabla clasificatoria.")
	fmt.Printf("\nLa puntuación máxima son %s\n\n", color.GreenString("15 puntos"))
}

func ShowGoodbyeScreen() {
	fmt.Println("\nFIN DEL JUEGO CAPITALES")
	color.Green(FinalArt)
}

//COLORES

func PrintSuccess(message string) {
	color.Green(message)
}

func PrintError(message string) {
	color.Red(message)
}

func PrintQuestion(message string) {
	color.Cyan(message)
}

func PrintLeaderboard(scores []game.PlayerScore) {
	fmt.Println("\n--- TABLA CLASIFICATORIA ---")

	color.Yellow("%-4s | %-20s | %s", "ID", "JUGADOR", "PUNTOS")
	fmt.Println("----------------------------------------")

	for _, score := range scores {
		fmt.Printf("%-4d | %-20s | %d\n", score.ID, score.Name, score.Points)
	}
	fmt.Println("----------------------------------------")
}

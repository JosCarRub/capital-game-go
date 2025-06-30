package tui

import "github.com/fatih/color"

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

package game

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"

	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func LoadCountries(filepath string) ([]Country, error) {
	fileBytes, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var countries []Country
	err = json.Unmarshal(fileBytes, &countries)
	if err != nil {
		return nil, err
	}

	return countries, nil
}

type Game struct {
	Countries        []Country
	RemainingIndices []int
	Hits             int
	Misses           int
	RoundSize        int
}

func NewGame(countries []Country) *Game {
	// lista de índices, desde 0 hasta el número total
	indices := make([]int, len(countries))
	for i := range countries {
		indices[i] = i
	}

	rand.New(rand.NewSource(time.Now().UnixNano()))

	return &Game{
		Countries:        countries,
		RemainingIndices: indices,
		Hits:             0,
		Misses:           0,
		RoundSize:        15,
	}

}

func (g *Game) NextQuestion() (*Country, bool) {

	if len(g.RemainingIndices) == 0 {
		return nil, false
	}

	randomIndexPosition := rand.Intn(len(g.RemainingIndices))
	// índice real del país en lista
	countryIndex := g.RemainingIndices[randomIndexPosition]

	// se elimina el índice para que no se repita
	g.RemainingIndices[randomIndexPosition] = g.RemainingIndices[len(g.RemainingIndices)-1]
	g.RemainingIndices = g.RemainingIndices[:len(g.RemainingIndices)-1]

	return &g.Countries[countryIndex], true
}

func NormalizeString(s string) string {
	// transformador para quitar los acentos.
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)

	result, _, err := transform.String(t, s)
	if err != nil {

		log.Printf("Error al normalizar string: %v", err)
		result = s
	}

	return strings.ToLower(strings.TrimSpace(result))
}

// comparador
func CheckAnswer(userInput string, correctCapital string) bool {
	normalizedUserInput := NormalizeString(userInput)
	normalizedCorrectCapital := NormalizeString(correctCapital)

	return normalizedUserInput == normalizedCorrectCapital
}

// incrementa aciertos
func (g *Game) RecordHit() {
	g.Hits++
}

// incrementa fallos
func (g *Game) RecordMiss() {
	g.Misses++
}

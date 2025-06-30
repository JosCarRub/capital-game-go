package game

import (
	"encoding/json"
	"os"
)

func LoadCountries(filepath string) ([]Country, error) {

	fileBytes, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var countries []Country
	// Decodifica (unmarshal)
	err = json.Unmarshal(fileBytes, &countries)
	if err != nil {
		return nil, err
	}

	return countries, nil
}

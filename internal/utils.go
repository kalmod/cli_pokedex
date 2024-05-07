package internal

import (
	"encoding/json"
	"fmt"
)

func (data *LocationAreaS) PrintAreas() {
	for _, location := range data.Results {
		fmt.Println(location.Name)
	}
}

func (data *LocationInformation) PrintPokemon() {
  for _, PokemonEncounter := range data.PokemonEncounters {
    fmt.Println(PokemonEncounter.Pokemon.Name)
  }
}

func PrintLocationNames(locationData *LocationAreaS) {
	for _, location := range locationData.Results {
		fmt.Println(location.Name)
	}

}

func PokeParseLocationAreaToJSON(locationData *LocationAreaS, val []byte) error {
	jsonUnmarshall_err := json.Unmarshal(val, &locationData)
	if jsonUnmarshall_err != nil {
		fmt.Println(jsonUnmarshall_err)
		return jsonUnmarshall_err
	}
	return  nil
}

func PokeParseLocationInfoToJSON(data *LocationInformation, val []byte)  error {
	jsonUnmarshall_err := json.Unmarshal(val, &data)
	if jsonUnmarshall_err != nil {
		fmt.Println(jsonUnmarshall_err)
		return jsonUnmarshall_err
	}
	return nil
}



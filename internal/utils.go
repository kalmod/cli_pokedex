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
  fmt.Println("Exploring " + data.Name + "...")
  fmt.Println("Found Pokemon:")
  for _, PokemonEncounter := range data.PokemonEncounters {
    fmt.Println(" - " + PokemonEncounter.Pokemon.Name)
  }
}

func (data *PokemonInfo) PrintPokemonInfo(){
  var allData string
  allData += fmt.Sprintf("Name: %v\n", data.Name)
  allData += fmt.Sprintf("Height: %v\n", data.Height)
  allData += fmt.Sprintf("Weight: %v\n", data.Weight)
  allData += fmt.Sprint("Stats:\n")
  for _, statInfo := range data.Stats {
    allData += fmt.Sprintf(" -%v: %v\n", statInfo.Stat.Name, statInfo.BaseStat)
  }
  allData += fmt.Sprint("Types:\n")
  for _, typeInfo := range data.Types {
    allData += fmt.Sprintf(" - %v\n",typeInfo.Type.Name)
  }
  fmt.Println(allData)
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


func PokeParsePokemonInfoToJSON(data *PokemonInfo, val []byte)  error {
	jsonUnmarshall_err := json.Unmarshal(val, &data)
	if jsonUnmarshall_err != nil {
		fmt.Println(jsonUnmarshall_err)
		return jsonUnmarshall_err
	}
	return nil
}

package main

import (
	"errors"
	"github.com/kalmod/cli_pokedex/internal"
)

func commandExplore(cfg *config, params ...string) error {
	if len(params) < 1 {
		return errors.New("No location given")
	}
  parsedLocationData := internal.LocationInformation{}

  // We check if the data for the chosen city is cached
  // Process stored data if it is.
  val, exists := cfg.cachedData.Get(params[0]); if exists {
    err := internal.PokeParseLocationInfoToJSON(&parsedLocationData, val) 
    if err != nil {
      return err
    }
    
  } else {
    // Else, perform request to retrieve data.
	  val, err := internal.ExploreTargetArea(params[0])
    if err != nil {
        return err
      }
    err = internal.PokeParseLocationInfoToJSON(&parsedLocationData, val) 
    if err != nil {
      return err
    }
    cfg.cachedData.Add(params[0],val)
  }

	// fmt.Println(val.PokemonEncounters)
  // I think key can be city name & 
  parsedLocationData.PrintPokemon()
  
	return nil
}

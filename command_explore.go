package main

import (
	"errors"
	"fmt"
	"github.com/kalmod/cli_pokedex/internal"
)

func commandExplore(cfg *config, params ...string) error {
	if len(params) > 1 {
		return errors.New("Too many params.")
	}
	val, err := internal.ExploreTargetArea(params[0])
	if err != nil {
		return err
	}

	// fmt.Println(val.PokemonEncounters)
  for _, PokemonEncounter := range val.PokemonEncounters {
    fmt.Println(PokemonEncounter.Pokemon.Name)
  }

	return nil
}

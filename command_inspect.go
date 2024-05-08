package main

import (
	"fmt"
)

func commandInspect(cfg *config, params ...string) error {
  if len(params) < 1 {
    fmt.Println("Pokemon name not given")
    return nil
  }
	pokemonInfo, pokedexExists := cfg.pokedex[params[0]]
	if pokedexExists {
		pokemonInfo.PrintPokemonInfo()
		return nil
	} else {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
}

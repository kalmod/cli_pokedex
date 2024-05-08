package main

import (
	"fmt"
	"github.com/kalmod/cli_pokedex/internal"
  "math/rand"
)

func commandCatch(cfg *config, params ...string) error {
  if  len(params) < 1 {
    fmt.Println("No name given")
    return nil
  }
  
  _, inPokedex := cfg.pokedex[params[0]]; if inPokedex {
    fmt.Printf("%v entry already in Pokedex\n", params[0])
    return nil
  }

	val, err := internal.CatchPokemon(params[0])
	if err != nil {
		return err
	}

	pokemonData := internal.PokemonInfo{}
	json_err := internal.PokeParsePokemonInfoToJSON(&pokemonData, val)
  if json_err != nil {
    return json_err
  }

  fmt.Printf("Throwing a Pokeball at %v...\n", params[0])
  catchPercentage := float32(rand.Intn(pokemonData.BaseExperience))/float32(pokemonData.BaseExperience)*100

  if catchPercentage < 50 {
    fmt.Printf("%v escaped!\n", params[0])
  } else {
    fmt.Printf("%v was caught!\n",params[0])
    fmt.Println("You may now inspect it with the \033[31minspect\033[0m command.")
    cfg.pokedex[params[0]] = pokemonData
  }


	return nil
}

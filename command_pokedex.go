package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, params ...string) error {
  if len(params) != 0 {
    return errors.New("Command takes no arguments..")
  }
  fmt.Println("Your Pokedex:")
  for name := range cfg.pokedex {
    fmt.Printf(" - \033[3;34m%v\033[0m\n", name )
  }
  return nil
}

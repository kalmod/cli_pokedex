package main

import "fmt"

func commandHelp(cfg *config) error {

  var Reset = "\033[0m"
  var Yellow = "\033[36m"

	fmt.Println(Yellow+"Welcome to the Pokedex"+Reset)
	fmt.Println(Yellow+"Usage:"+Reset)
  fmt.Println()
	commands := getCommand()
	for _,command := range commands {
		fmt.Printf("\033[31m%s\033[0m: %s\n", command.name, command.help)
	}
	return nil
}

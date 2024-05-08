package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/kalmod/cli_pokedex/internal"
)

type cliCommand struct {
	name     string
	help     string
	callback func(*config, ...string) error
}

type config struct {
	Next       *string
	Previous   *string
	cachedData internal.Cache
	pokedex    map[string]internal.PokemonInfo
}

func getCommand() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:     "help",
			help:     "Displays a help message",
			callback: commandHelp,
		},
		"exit": {
			name:     "exit",
			help:     "Exit the pokedex",
			callback: commandExit,
		},
		"map": {
			name:     "map",
			help:     "Show the next 20 location areas",
			callback: commandMap,
		},
		"mapb": {
			name:     "mapb",
			help:     "Show the previous 20 location areas",
			callback: commandMapBack,
		},
		"explore": {
			name:     "explore [Area Name]",
			help:     "Show the pokemon that can be found in the specified area",
			callback: commandExplore,
		},
		"catch": {
			name:     "catch [Pokemon Name]",
			help:     "Use a pokemon name with the command to catch the pokemon",
			callback: commandCatch,
		},
		"inspect": {
			name:     "inspect [Pokemon Name]",
			help:     "Use a pokemon name with the command to get pokemon info",
			callback: commandInspect,
		},
		"pokedex": {
			name:     "pokedex",
			help:     "Display all entries in the pokedex",
			callback: commandPokedex,
		},
	}
}

func cleanInput(text string) []string {
	formattedText := strings.ToLower(text)
	words := strings.Fields(formattedText)
	if len(words) == 0 {
		return []string{" "}
	}
	return words
}

func repl() {
	interval := time.Second * 5
	cfg := config{cachedData: internal.NewCache(interval), pokedex: map[string]internal.PokemonInfo{}}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\033[33mpokedex\033[0m > ")
		scanner.Scan()
		allWords := cleanInput(scanner.Text())
		command := allWords[0]
		params := allWords[1:]

		if call, ok := getCommand()[command]; ok {
			err := call.callback(&cfg, params...)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

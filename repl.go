package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
  "github.com/kalmod/cli_pokedex/internal"
)

type cliCommand struct {
	name     string
	help     string
	callback func(*config) error
}

type config struct {
  Next *string
  Previous *string
  cachedData internal.Cache
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
      name: "map",
      help: "Show the next 20 location areas",
      callback: commandMap,
    },
    "mapb": {
      name: "mapb",
      help: "Show the previous 20 location areas",
      callback: commandMapBack,
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
  cfg := config{cachedData: internal.NewCache(5)}
	scanner := bufio.NewScanner(os.Stdin)
	for {
	  fmt.Print("\033[33mpokedex\033[0m > ")
    scanner.Scan() 
		allWords := cleanInput(scanner.Text())
		command := allWords[0]

		if call, ok := getCommand()[command]; ok {
      err := call.callback(&cfg)
      if err != nil {
        fmt.Println(err)
      }
		} else {
      fmt.Println("Unknown command")
      continue
    }
	}
}

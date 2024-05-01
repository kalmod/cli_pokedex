package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var locationData = locationAreaS{}

type cliCommand struct {
	name     string
	help     string
	callback func() error
}

type locationAreaS struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
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
      callback: func ()error {
        err := commandMap(&locationData)
        return err
      },
    },
	}
}

func cleanInput(text string) []string {
	formattedText := strings.ToLower(text)
	words := strings.Fields(formattedText)
	return words
}

func repl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
	  fmt.Print("\033[33mpokedex\033[0m > ")
    scanner.Scan() 
		allWords := cleanInput(scanner.Text())
		command := allWords[0]

		if call, ok := getCommand()[command]; ok {
      err := call.callback()
      if err != nil {
        fmt.Println(err)
      }
		} else {
      fmt.Println("Unknown command")
      continue
    }
	}
}

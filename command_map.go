package main

import (
	"fmt"

	"github.com/kalmod/cli_pokedex/internal"
)

func commandMap(cfg *config, params ...string) error {
	url := internal.BaseUrl + "/location-area/"
	if cfg.Next != nil {
		url = *cfg.Next
	}

	locationData := internal.LocationAreaS{}
	var err error = nil

	if val, exists := cfg.cachedData.Get(url); exists {
		err = internal.PokeParseLocationAreaToJSON(&locationData, val)
		if err != nil {
			return err
		}
	} else {
		val, err := internal.GetLocationAreas(url)
		cfg.cachedData.Add(url, val)
		if err != nil {
			return err
		}
		err = internal.PokeParseLocationAreaToJSON(&locationData, val)
		if err != nil {
			return err
		}
	}

	cfg.Next = locationData.Next
	cfg.Previous = locationData.Previous

	locationData.PrintAreas()

	return nil
}

func commandMapBack(cfg *config, params ...string) error {

	if cfg.Previous == nil {
		fmt.Println("ERROR: No Previous url")
		return nil
	}
	url := *cfg.Previous

	locationData := internal.LocationAreaS{}
	var err error = nil

	if val, exists := cfg.cachedData.Get(url); exists {
		err = internal.PokeParseLocationAreaToJSON(&locationData, val)
		if err != nil {
			return err
		}

	} else {
		val, err := internal.GetLocationAreas(url)
		if err != nil {
			return err
		}
		cfg.cachedData.Add(url, val)
		err = internal.PokeParseLocationAreaToJSON(&locationData, val)
		if err != nil {
			return err
		}
	}

	cfg.Next = locationData.Next
	cfg.Previous = locationData.Previous

	locationData.PrintAreas()

	return nil
}

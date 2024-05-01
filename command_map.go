package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func commandMap(locationData *locationAreaS) error {
	startUrl := "https://pokeapi.co/api/v2/location-area/"
	var nextUrl = ""
	if locationData.Next == nil {
		nextUrl = startUrl
	} else {
		nextUrl = *locationData.Next
	}
	res, err := http.Get(nextUrl)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf(
			"Response failed with status code: %d and \nbody: %s\n",
			res.StatusCode, body,
		)
	}
	if err != nil {
		log.Fatal(err)
	}
	jsonUnmarshall_err := json.Unmarshal(body, &locationData)
	if jsonUnmarshall_err != nil {
		fmt.Println(err)
	}
	fmt.Println(locationData.Count)
	fmt.Println(*locationData.Next)

	if locationData.Previous != nil {
		fmt.Println(*locationData.Previous)
	}
  for _, location := range locationData.Results {
    fmt.Println(location.Name)
  }
	return nil
}

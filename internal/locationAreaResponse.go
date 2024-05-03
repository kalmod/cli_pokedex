package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type LocationAreaS struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreas(targetUrl string) ([]byte, error) {


	res, err := http.Get(targetUrl)

	if err != nil {
		log.Fatal(err)
		return []byte{}, err
	}

	if res.StatusCode > 299 {
		log.Fatalf(
			"Response failed with status code: %d \n",
			res.StatusCode,
		)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	// TODO Our cached data would replay what's in the body.
	// io.ReadAll returns a []byte

	if err != nil {
		log.Fatal(err)
		return []byte{}, err
	}

  return body, nil
}

func PrintLocationNames(locationData *LocationAreaS) {
	for _, location := range locationData.Results {
		fmt.Println(location.Name)
	}

}

func PokeParseToJSON(locationData *LocationAreaS, val []byte) (LocationAreaS, error) {
	jsonUnmarshall_err := json.Unmarshal(val, &locationData)
	if jsonUnmarshall_err != nil {
		fmt.Println(jsonUnmarshall_err)
		return *locationData, jsonUnmarshall_err
	}
  return *locationData, nil
}


// func GetLocationAreas(targetUrl string) (LocationAreaS, error) {

// 	locationData := LocationAreaS{}
// 	res, err := http.Get(targetUrl)

// 	if err != nil {
// 		log.Fatal(err)
// 		return locationData, err
// 	}

// 	if res.StatusCode > 299 {
// 		log.Fatalf(
// 			"Response failed with status code: %d \n",
// 			res.StatusCode,
// 		)
// 	}
// 	defer res.Body.Close()

// 	body, err := io.ReadAll(res.Body)
// 	// TODO Our cached data would replay what's in the body.
// 	// io.ReadAll returns a []byte

// 	if err != nil {
// 		log.Fatal(err)
// 		return locationData, err
// 	}

//   return PokeParseToJSON(&locationData, body)
// }

// func GetPreviousLocationAreas(targetUrl string) (LocationAreaS, error) {
//
//   locationData := LocationAreaS{}
// 	res, err := http.Get(targetUrl)

// 	if err != nil {
// 		log.Fatal(err)
//     return locationData, err
// 	}

// 	body, err := io.ReadAll(res.Body)
// 	defer res.Body.Close()

// 	if res.StatusCode > 299 {
// 		log.Fatalf(
// 			"Response failed with status code: %d and \nbody: %s\n",
// 			res.StatusCode, body,
// 		)
// 	}

// 	if err != nil {
// 		log.Fatal(err)
//     return locationData, err
// 	}
//
// 	jsonUnmarshall_err := json.Unmarshal(body, &locationData)
//
// 	if jsonUnmarshall_err != nil {
// 		fmt.Println(err)
//     return locationData, err
// 	}

//     return locationData, nil
// }

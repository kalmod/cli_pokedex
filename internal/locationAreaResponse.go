package internal

import (
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

	if err != nil {
		log.Fatal(err)
		return []byte{}, err
	}

  return body, nil
}

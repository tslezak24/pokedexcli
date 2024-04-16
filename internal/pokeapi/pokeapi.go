package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Response struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Location struct {
	Name string
	URL  string
}

func MapLocations() []Location {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		fmt.Print(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		fmt.Printf("Response failed with status code: %v", res.StatusCode)
	}
	if err != nil {
		fmt.Print(err)
	}
	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}
	locations := make([]Location, 0)
	for _, location := range response.Results {
		locations = append(locations, Location{Name: location.Name, URL: location.URL})
	}
	return locations
}

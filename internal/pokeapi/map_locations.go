package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func MapLocations(locationsURL string) Response {
	if locationsURL == "" {
		locationsURL = url
	}
	res, err := http.Get(locationsURL)
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

	return response
}

package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) CatchPokemon(pokemonName string) Pokemon {
	catchURL := url + "pokemon/" + pokemonName

	res, err := http.Get(catchURL)
	if err != nil {
		fmt.Println("Error fetching data")
	}
	if res.StatusCode > 299 {
		fmt.Printf("error: %v", res.Status)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error")
	}

	var response Pokemon
	if json.Unmarshal(body, &response) != nil {
		fmt.Println("Unable to unmarshal JSON")
	}

	return response
}

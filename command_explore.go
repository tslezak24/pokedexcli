package main

import (
	"fmt"
)

func commandExplore(config *config, area string) error {
	res := config.pokeapiClient.ExploreArea(area)
	for _, pokemon := range res.PokemonEncounters {
		fmt.Printf(" - %v \n", pokemon.Pokemon.Name)
	}
	return nil
}

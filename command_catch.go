package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(config *config, pokemonName string) error {
	res := config.pokeapiClient.CatchPokemon(pokemonName)

	baseExp := res.BaseExperience / 8
	randomNum := rand.Intn(baseExp + 1)

	if baseExp == randomNum {
		fmt.Printf("%v was caught! \n", res.Name)
		config.user.pokedex[res.Name] = res
	} else {
		fmt.Printf("%v escaped! \n", res.Name)
	}

	return nil
}

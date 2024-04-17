package main

import (
	"fmt"
	"internal/pokeapi"
)

func commandMap(config *config) error {
	res := pokeapi.MapLocations(config.next)

	for _, location := range res.Results {
		fmt.Println(location.Name)
	}
	config.next = res.Next
	config.previous = res.Previous
	return nil
}

func commandMapBack(config *config) error {
	res := pokeapi.MapLocations(config.previous)

	for _, location := range res.Results {
		fmt.Println(location.Name)
	}
	config.previous = res.Previous
	config.previous = res.Previous
	return nil
}

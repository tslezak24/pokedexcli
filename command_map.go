package main

import (
	"fmt"
)

func commandMap(config *config, area string) error {
	res := config.pokeapiClient.MapLocations(config.next)

	for _, location := range res.Results {
		fmt.Println(location.Name)
	}
	config.next = res.Next
	config.previous = res.Previous
	return nil
}

func commandMapBack(config *config, area string) error {
	res := config.pokeapiClient.MapLocations(config.previous)

	for _, location := range res.Results {
		fmt.Println(location.Name)
	}
	config.previous = res.Previous
	config.previous = res.Previous
	return nil
}

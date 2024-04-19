package main

import (
	"fmt"
)

func commandPokedex(config *config, area string) error {
	if len(config.user.pokedex) == 0 {
		fmt.Println("You have no pokemon yet")
	} else {
		fmt.Println("Your Pokedex:")
		for pokemon := range config.user.pokedex {
			fmt.Printf("  - %v \n", pokemon)
		}
	}
	return nil
}

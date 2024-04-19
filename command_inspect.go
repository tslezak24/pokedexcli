package main

import "fmt"

func commandInspect(config *config, pokemon string) error {
	if poke, ok := config.user.pokedex[pokemon]; ok {
		fmt.Printf("Name: %v \nHeight: %v \nWeight: %v \nStats: \n", poke.Name, poke.Height, poke.Weight)
		for _, stat := range poke.Stats {
			fmt.Printf("    -%v: %v \n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, t := range poke.Types {
			fmt.Printf("    - %v \n", t.Type.Name)
		}
	} else {
		fmt.Printf("you have not caught %v \n", pokemon)
	}
	return nil
}

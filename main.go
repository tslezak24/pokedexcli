package main

import (
	"internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	usr := &User{
		pokedex: make(map[string]pokeapi.Pokemon),
	}
	cfg := &config{
		pokeapiClient: pokeClient,
		user:          *usr,
	}
	start(cfg)
}

package main

import (
	"bufio"
	"fmt"
	"internal/pokeapi"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

type config struct {
	pokeapiClient pokeapi.Client
	previous      string
	next          string
}

type Location struct {
	Name string
	URL  string
}

func start(cfg *config) {
	commands := defineCommands()
	cliName := "pokedex"
	reader := bufio.NewScanner(os.Stdin)
	printPrompt(cliName)
	for reader.Scan() {
		inputs := strings.Fields(reader.Text())
		comm := inputs[0]
		area := ""
		if len(inputs) > 1 {
			area = inputs[1]
		}
		if command, exists := commands[comm]; exists {
			command.callback(cfg, area)
		} else {
			printUnkown(cliName)
		}
		printPrompt(cliName)
	}
}

func printPrompt(cliName string) {
	fmt.Print(cliName, "> ")
}

func printUnkown(text string) {
	fmt.Println(text, ": command not found")
}

func defineCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "Find pokemon in area",
			callback:    commandExplore,
		},
	}
}

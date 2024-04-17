package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	previous string
	next     string
}

type Location struct {
	Name string
	URL  string
}

func start() {
	commands := defineCommands()
	cliName := "pokedex"
	reader := bufio.NewScanner(os.Stdin)
	printPrompt(cliName)
	cfg := &config{}
	for reader.Scan() {
		text := reader.Text()
		if command, exists := commands[text]; exists {
			command.callback(cfg)
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
	}
}

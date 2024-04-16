package main

import (
	"bufio"
	"fmt"
	"internal/pokeapi"
	"os"
	"os/exec"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	commands := defineCommands()
	cliName := "pokedex"
	reader := bufio.NewScanner(os.Stdin)
	printPrompt(cliName)
	for reader.Scan() {
		text := reader.Text()
		if command, exists := commands[text]; exists {
			command.callback()
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

func commandHelp() error {
	fmt.Println("Welcome. Here are some useful commands:")
	return nil
}

func clearScreen() error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	return nil
}

func commandMap() error {
	res := pokeapi.MapLocations()
	for _, r := range res {
		fmt.Println(r.Name)
	}
	return nil
}

func commandMapBack() error {
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
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
		"clear": {
			name:        "clear",
			description: "Clear Screen",
			callback:    clearScreen,
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

package commands

import (
	"goPokedex/repl"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*repl.config, ...string) error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message including a list of available commands",
			callback:    CommandHelp,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Attempts to catch the entered Pokemon",
			callback:    CommandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "View details about a caught Pokemon",
			callback:    CommandInspect,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explore the entered location",
			callback:    CommandExplore,
		},
		"mapnext": {
			name:        "map next",
			description: "View the next page of locations",
			callback:    CommandMapf,
		},
		"mapprev": {
			name:        "map previous",
			description: "View the previous page of locations",
			callback:    CommandMapb,
		},
		"pokedex": {
			name:        "pokedex",
			description: "View all of your caught Pokemon",
			callback:    CommandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    CommandExit,
		},
	}
}

package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/DavidIRL/goPokedex/internal/pokeapi"
)

type Config struct {
	PokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	CaughtPokemon    map[string]pokeapi.Pokemon
}

func StartRepl(conf *Config) {
    scanner := bufio.NewScanner(os.Stdin)
	for { //ever
		fmt.Print("goPokedex > ")
		scanner.Scan()

		text := cleanInput(scanner.Text())
		if len(text) == 0 {
			continue
		}
		commandName := text[0]
		args := []string{}
		if len(text) > 1 {
			args = text[1:]
		}

		command, exists := GetCommands()[commandName]
		if exists {
			err := command.callback(conf, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command entered\nFor list of commands, enter 'help'")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, ...string) error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message including a list of available commands",
			callback:    commandHelp,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Attempts to catch the entered Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "View details about a caught Pokemon",
			callback:    commandInspect,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explore the entered location",
			callback:    commandExplore,
		},
		"map": {
			name:        "map",
			description: "View the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "View the previous page of locations",
			callback:    commandMapb,
		},
		"pokedex": {
			name:        "pokedex",
			description: "View all of your caught Pokemon",
			callback:    commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

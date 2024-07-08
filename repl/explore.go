package repl

import (
	"errors"
	"fmt"
)

func commandExplore(conf *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("A location name must be provided")
	}

	name := args[0]
	location, err := conf.PokeapiClient.GetLocation(name)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Printf("Found Pokemon: ")
	for _, fight := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", fight.Pokemon.Name)
	}
	return nil
}

package repl

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(conf *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("A Pokemon name must be provided")
	}

	name := args[0]
	pokemon, err := conf.PokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	res := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if res > 45 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)
	fmt.Printf("\nYou can now inspect %s with the inspect command.\n", pokemon.Name)

	conf.CaughtPokemon[pokemon.Name] = pokemon
	return nil
}

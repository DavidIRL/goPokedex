package commands

import (
	"errors"
	"fmt"
	repl "goPokedex/repl"
	"math/rand"
)

func CommandCatch(conf *repl.config, args ...string) error {
	if len(args) != 1 {
		return errors.New("A Pokemon name must be provided")
	}

	name := args[0]
	pokemon, error := conf.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	res := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.name)
	if res > 45 {
		fmt.Printf("%s escaped!\n", pokemon.name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.name)
	fmt.Printf("\nYou can now inspect %s with the inspect command.", pokemon.name)

	conf.caughtPokemon[pokemon.name] = pokemon
	return nil
}

package commands

import (
	"fmt"
	"goPokedex/repl"
)

func CommandPokedex(conf *repl.config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range conf.caughtPokemon {
		fmt.Printf(" -%s\n", pokemon.Name)
	}
	return nil
}

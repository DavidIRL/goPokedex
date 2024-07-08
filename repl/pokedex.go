package repl

import (
	"fmt"
)

func commandPokedex(conf *Config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range conf.CaughtPokemon {
		fmt.Printf(" -%s\n", pokemon.Name)
	}
	return nil
}

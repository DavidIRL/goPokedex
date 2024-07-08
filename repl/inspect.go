package repl

import (
	"errors"
	"fmt"
)

func commandInspect(conf *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("A Pokemon name must be provided")
	}

	name := args[0]
	pokemon, ok := conf.CaughtPokemon[name]
	if !ok {
		return errors.New("You have not yet caught that Pokemon")
	}

	fmt.Println("Name: ", pokemon.Name)
	fmt.Println("Height: ", pokemon.Height)
	fmt.Println("Weight: ", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeInfo := range pokemon.Types {
		fmt.Println(" -", typeInfo.Type.Name)
	}
	return nil

}

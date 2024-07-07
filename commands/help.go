package commands

import (
	"fmt"
	"goPokedex/repl"
)

func CommandHelp(conf *repl.config, args ...string) error {
	fmt.Println("\nWelcome to the goPokedex!")
	fmt.Println("Usage:\n")
	for _, cmd := range GetCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	// make space for next input session
	fmt.Println()
	return nil
}

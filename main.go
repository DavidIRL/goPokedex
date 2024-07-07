package main

import (
	"github.com/bootdotdev/go-api-gate/courses/projects/pokedexcli/internal/pokeapi"
	command "goPokedex/commands"
	repl "goPokedex/repl"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	conf := &repl.config{
		pokeapiClient: pokeClient,
		caughtPokemon: map[string]pokeapi.Pokemon{},
	}

	repl.StartRepl(conf)
}

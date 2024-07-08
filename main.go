package main

import (
	"github.com/DavidIRL/goPokedex/internal/pokeapi"
	"github.com/DavidIRL/goPokedex/repl"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	conf := &repl.Config{
		PokeapiClient: pokeClient,
		CaughtPokemon: map[string]pokeapi.Pokemon{},
	}

	repl.StartRepl(conf)
}

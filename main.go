package main

import (
    "time"

    "github.com/bootdobdev/go-api-gate/courses/projects/pokedexcli/internal/pokeapi"
)


func main() {
    pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
    conf := &config{
        pokeapiClient: pokeClient,
        caughtPokemon: map[string]pokeapi.Pokemon{},
    }

    startRepl(conf)
}

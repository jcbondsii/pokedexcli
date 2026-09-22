package main

import (
	"time"

	"pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	caughtPokemon := make(map[string]pokeapi.PokemonInfo)
	cfg := &config{
		commands:      getCommands(),
		pokeapiClient: pokeClient,
		caughtPokemon: caughtPokemon,
	}

	startRepl(cfg)
}

package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: catch <pokemon_name>")
	}
	pokemonName := args[0]
	pokemonInfo, err := cfg.pokeapiClient.FetchPokemonInfo(pokemonName)
	if err != nil {
		return fmt.Errorf("failed to fetch pokemon info: %v", err)
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	// Implementation for catching a pokemon
	catchRate := float64(pokemonInfo.BaseExperience) / 100.0 // Example catch rate based on base experience
	isCaught := rand.Float64() < catchRate
	if isCaught {
		fmt.Printf(" %s was caught!\n", pokemonName)
		cfg.caughtPokemon[pokemonName] = pokemonInfo
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}
	return nil
}

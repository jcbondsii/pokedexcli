package main

import (
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: inspect <pokemon_name>")
	}
	pokemonName := args[0]
	if _, exists := cfg.caughtPokemon[pokemonName]; !exists {
		return fmt.Errorf("you have not caught that pokemon yet")
	}
	for _, caughtPokemon := range cfg.caughtPokemon {
		if caughtPokemon.Name == pokemonName {
			fmt.Printf("Name: %s\n", caughtPokemon.Name)
			fmt.Printf("Height: %d\n", caughtPokemon.Height)
			fmt.Printf("Weight: %d\n", caughtPokemon.Weight)
			fmt.Println("Stats: ")
			for _, stat := range caughtPokemon.Stats {
				fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
			}
			fmt.Println("Types: ")
			for _, typeInfo := range caughtPokemon.Types {
				fmt.Printf("  - %s\n", typeInfo.Type.Name)
			}
			return nil
		}
	}
	return nil
}

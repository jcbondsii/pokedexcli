package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(args) != 0 {
		return fmt.Errorf("usage: pokedex")
	}
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("You have not caught any Pokemon yet.")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for name := range cfg.caughtPokemon {
		fmt.Println(" - " + name)
	}
	return nil
}

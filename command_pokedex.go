package main

import (
	"fmt"
)

func commandPokedex(cfg *config, pokemonName string) error {
	fmt.Println("Your Pokedex:")
	for _, v := range cfg.pokedex {
		fmt.Printf("- %v\n", v.Name)
	}

	return nil
}

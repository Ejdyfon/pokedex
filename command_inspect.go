package main

import (
	"fmt"
)

func commandInspect(cfg *config, pokemonName string) error {
	pokemon, exists := cfg.pokedex[pokemonName]
	if exists {
		fmt.Println("Name: " + pokemon.Name)
		fmt.Printf("Height: %v", pokemon.Height)
		fmt.Printf("Weight: %v\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, v := range pokemon.Stats {
			fmt.Printf("   -%v:  %v\n", v.Stat.Name, v.BaseStat)
		}
		fmt.Println("Types:")
		for _, v := range pokemon.Types {
			fmt.Printf("   - %v\n", v.Type.Name)
		}
	} else {
		fmt.Println("You dont have " + pokemonName)
	}

	return nil
}

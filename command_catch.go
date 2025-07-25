package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, pokemonName string) error {
	fmt.Println("Throwing a Pokeball at " + pokemonName + "...")
	locationsResp, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	chance := locationsResp.BaseExperience
	throw := 30 + rand.Intn(100)

	if throw >= chance {
		cfg.pokedex[pokemonName] = locationsResp
		fmt.Println(pokemonName + " was caught!")
	} else {
		fmt.Println(pokemonName + " escaped!")
	}

	return nil
}

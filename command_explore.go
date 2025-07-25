package main

import (
	"fmt"
)

func commandExplore(cfg *config, areaName string) error {
	locationsResp, err := cfg.pokeapiClient.ListLocationsExtended(areaName)
	if err != nil {
		return err
	}

	for _, loc := range locationsResp.PokemonEncounters {
		fmt.Println(loc.Pokemon.Name)
	}
	return nil
}

package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// GetPokemon -
func (c *Client) GetPokemon(name string) (PokemonDetails, error) {
	url := baseURL + "/pokemon/" + name

	if val, ok := c.cache.Get(url); ok {
		locationsResp := PokemonDetails{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return PokemonDetails{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonDetails{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonDetails{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonDetails{}, err
	}

	locationsResp := PokemonDetails{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return PokemonDetails{}, err
	}

	c.cache.Add(url, dat)
	return locationsResp, nil
}

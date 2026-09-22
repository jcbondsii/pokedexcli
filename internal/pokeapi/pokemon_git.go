package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// Fetch PokemonInfo fetches the information of a specific pokemon by name
func (c *Client) FetchPokemonInfo(pokemonName string) (PokemonInfo, error) {
	url := baseURL + "/pokemon/" + pokemonName

	if val, ok := c.cache.Get(url); ok {
		pokemonInfo := PokemonInfo{}
		err := json.Unmarshal(val, &pokemonInfo)
		if err != nil {
			return PokemonInfo{}, err
		}

		return pokemonInfo, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonInfo{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonInfo{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonInfo{}, err
	}

	pokemonInfo := PokemonInfo{}
	err = json.Unmarshal(dat, &pokemonInfo)
	if err != nil {
		return PokemonInfo{}, err
	}

	c.cache.Add(url, dat)
	return pokemonInfo, nil
}

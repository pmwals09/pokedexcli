package pokeapi

import (
	"encoding/json"
  "fmt"
	"io"
	"net/http"
	"time"

	"github.com/pmwals09/pokedexcli/internal/pokecache"
)

var cache = pokecache.NewCache(time.Second)

func GetLocationAreas(next *string) (LocationAreasResponse, error) {
	apiResponse := LocationAreasResponse{}
	path := "https://pokeapi.co/api/v2/location-area?limit=20"
	if next != nil {
		path = *next
	}

  body, err := checkCache(path)
  if err != nil {
    return apiResponse, err
  }

	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return apiResponse, err
	}

	return apiResponse, nil
}

func GetLocationArea(locationName string) (LocationAreaResponse, error) {
	apiResponse := LocationAreaResponse{}
  path := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", locationName)

  body, err := checkCache(path)
  if err != nil {
    return apiResponse, err
  }

  err = json.Unmarshal(body, &apiResponse)
  if err != nil {
    return apiResponse, err
  }

  return apiResponse, nil
}

func GetPokemon(pokemonName string) (PokemonResponse, error) {
  apiResponse := PokemonResponse{}
  path := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemonName)

  body, err := checkCache(path)
  if err != nil {
    return apiResponse, err
  }

  err = json.Unmarshal(body, &apiResponse)
  if err != nil {
    return apiResponse, err
  }
  return apiResponse, nil
}

func checkCache(path string) ([]byte, error) {
  var body []byte
  if cacheEntry, ok := cache.Get(path); ok {
    return cacheEntry, nil
  }
  res, err := http.Get(path)
  if err != nil {
    return body, err
  }
  body, err = io.ReadAll(res.Body)
  defer res.Body.Close()
  if err != nil {
    return body, err
  }
  cache.Add(path, body)
  return body, nil
}

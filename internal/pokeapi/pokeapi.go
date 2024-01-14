package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
  "time"

  "github.com/pmwals09/pokedexcli/internal/pokecache"
)

var cache = pokecache.NewCache(time.Second)

type APIResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationArea(next *string) (APIResponse, error) {
  apiResponse := APIResponse{}
  path := "https://pokeapi.co/api/v2/location-area?limit=20"
  if next != nil {
    path = *next
  }

  var body []byte
  if cacheEntry, ok := cache.Get(path); ok {
    body = cacheEntry
  } else {
    res, err := http.Get(path)
    if err != nil {
      return apiResponse, err
    }
    body, err = io.ReadAll(res.Body)
    if err != nil {
      return apiResponse, err
    }
    cache.Add(path, body)
  }
  err := json.Unmarshal(body, &apiResponse)
  if err != nil {
    return apiResponse, err
  }
  return apiResponse, nil
}

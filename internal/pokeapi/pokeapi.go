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

type LocationAreasResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type LocationAreaResponse struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func GetLocationAreas(next *string) (LocationAreasResponse, error) {
	apiResponse := LocationAreasResponse{}
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

func GetLocationArea(locationName string) (LocationAreaResponse, error) {
	apiResponse := LocationAreaResponse{}
  path := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", locationName)
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

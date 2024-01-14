package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

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
	res, err := http.Get(path)
	if err != nil {
		return apiResponse, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return apiResponse, err
	}
  err = json.Unmarshal(body, &apiResponse)
  if err != nil {
    return apiResponse, err
  }
  return apiResponse, nil
}

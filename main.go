package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/pmwals09/pokedexcli/internal/pokeapi"
)

const prompt = "pokedex > "

type config struct {
	locationNext *string
	locationPrev *string
}

type command struct {
	name        string
	description string
	callback    func([]string, *config) error
}

var commands = map[string]command{
	"exit": {
		name:        "exit",
		description: "Exits the program",
		callback:    exitCb,
	},
	"map": {
		name:        "map",
		description: "Displays the names of 20 location areas in the Pokemon world",
		callback:    mapCb,
	},
	"mapb": {
		name:        "mapb",
		description: "Displays the previous 20 location areas in the Pokemon world",
		callback:    mapbCb,
	},
  "explore": {
    name: "explore",
    description: "See a list of all the Pokémon in a given area",
    callback: exploreCb,
  },
}

func exitCb(parameters []string, config *config) error {
	os.Exit(0)
	return nil
}

func mapCb(parameters []string, config *config) error {
  res, err := pokeapi.GetLocationAreas(config.locationNext)
  if err != nil {
    return err
  }
  config.locationNext = res.Next
  config.locationPrev = res.Previous
  for _, loc := range res.Results {
    fmt.Println(loc.Name)
  }
  return nil
}

func mapbCb(parameters []string, config *config) error {
  res, err := pokeapi.GetLocationAreas(config.locationPrev)
  if err != nil {
    return err
  }
  config.locationNext = res.Next
  config.locationPrev = res.Previous
  for _, loc := range res.Results {
    fmt.Println(loc.Name)
  }
  return nil
}

func exploreCb(parameters []string, config *config) error {
  fmt.Printf("Exploring %s...\n", parameters[0])
  res, err := pokeapi.GetLocationArea(parameters[0])
  if err != nil {
    return err
  }
  fmt.Println("Found the following Pokémon:")
  for _, encounter := range res.PokemonEncounters {
    fmt.Printf("- %s\n", encounter.Pokemon.Name)
  }
  return nil
}

var helpCommand = command{
	name:        "help",
	description: "Displays this help message",
	callback: func(parameters []string, config *config) error {
		for _, cmd := range commands {
			fmt.Printf("%s - %s\n", cmd.name, cmd.description)
		}
		return nil
	},
}

func main() {
	commands["help"] = helpCommand
	config := &config{}

	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))

Loop:
	for {
		fmt.Print(prompt)
		for scanner.Scan() {
			text := scanner.Text()
      tokens := strings.Fields(text)
      cmdText, parameters := tokens[0], tokens[1:]
			if cmd, ok := commands[cmdText]; ok {
				err := cmd.callback(parameters, config)
				if err != nil {
					fmt.Printf("Error in command %s: %s\n", cmd.name, err.Error())
				}
			} else {
				fmt.Println("Unknown command")
				commands["help"].callback(parameters, config)
			}
			continue Loop
		}
	}
}

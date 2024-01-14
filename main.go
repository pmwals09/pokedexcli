package main

import (
	"bufio"
	"fmt"
	"os"

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
	callback    func(*config) error
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
}

func exitCb(config *config) error {
	os.Exit(0)
	return nil
}

func mapCb(config *config) error {
  res, err := pokeapi.GetLocationArea(config.locationNext)
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

func mapbCb(config *config) error {
  res, err := pokeapi.GetLocationArea(config.locationPrev)
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

var helpCommand = command{
	name:        "help",
	description: "Displays this help message",
	callback: func(config *config) error {
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
			cmdText := scanner.Text()
			if cmd, ok := commands[cmdText]; ok {
				err := cmd.callback(config)
				if err != nil {
					fmt.Printf("Error in command %s: %s\n", cmd.name, err.Error())
				}
			} else {
				fmt.Println("Unknown command")
				commands["help"].callback(config)
			}
			continue Loop
		}
	}
}

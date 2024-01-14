package main

import (
	"bufio"
	"fmt"
	"os"
)

const prompt = "pokedex > "

type command struct {
	name        string
	description string
	callback    func() error
}


func main() {
  commands := map[string]command{
		"exit": {
			name:        "exit",
			description: "Exits the program",
			callback:    exitCb,
		},
	}

  helpCommand := command {
    name: "help",
    description: "Displays this help message",
    callback: func() error {
      for _, cmd := range commands {
        fmt.Printf("%s - %s\n", cmd.name, cmd.description)
      }
      return nil
    },
  }

  commands["help"] = helpCommand

	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))

Loop:
	for {
		fmt.Print(prompt)
		for scanner.Scan() {
      cmdText := scanner.Text()
      if cmd, ok := commands[cmdText]; ok {
        cmd.callback()
      } else {
        fmt.Println("Unknown command")
        commands["help"].callback()
      }
			continue Loop
		}
	}
}

func exitCb() error {
  os.Exit(0)
  return nil
}

package main

import (
	"fmt"
)

var HELP = `Welcome to the Pokedex!
Usage:


`

func HelpAction(cfg *config) error {
	updateHelp(cfg.registry)
	fmt.Print(HELP)
	return nil
}

func updateHelp(registry map[string]cliCommand) {
	for _, command := range registry {
		HELP = HELP + fmt.Sprintf("%s: %s\n", command.name, command.description)
	}
}

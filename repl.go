package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	registry map[string]cliCommand
}

func startrepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Printf("ScannerError: %v", err)
			return
		}
		cmdtext := scanner.Text()
		cmdlist := cleanInput(cmdtext)

		command := strings.ToLower(cmdlist[0])
		value, ok := cfg.registry[command]
		if !ok {
			print("Unknown command")
			continue
		}

		err := value.callback(cfg)
		if err != nil {
			fmt.Printf("ActionError: %s", err)
		}
	}
}

func cleanInput(text string) []string {
	result := strings.Split(text, " ")
	for _, item := range result {
		if item == "" {
			result = slices.DeleteFunc(result, func(item string) bool {
				return item == ""
			})
		}
	}
	return result
}

package main

func main() {
	// fmt.Println("Hello, World!")

	var cfg config
	cfg.registry = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    HelpAction,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    ExitAction,
		},
		"quit": {
			name:        "quit",
			description: "Exit the Pokedex",
			callback:    ExitAction,
		},
	}
	startrepl(&cfg)
}

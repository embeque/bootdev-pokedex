package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/embeque/pokedexcli/internal/pokecache"
)

const pokeapi = "https://pokeapi.co/api/v2/"

type config struct {
	registry map[string]cliCommand
	Next     string
	Prev     string
	Cache    pokecache.Cache
	logger   *slog.Logger
}

func main() {
	// fmt.Println("Hello, World!")

	var cfg config
	cfg.registry = map[string]cliCommand{
		"map": {
			name:        "map",
			description: "List the Location Areas of Map",
			callback:    ActionNextMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List the previous Location Areas of Map",
			callback:    ActionPrevMap,
		},
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
	cfg.Next = pokeapi + "location-area?offset=0&limit=20"
	cfg.Cache = pokecache.NewCache(time.Minute * 5)

	// logger
	file, err := os.OpenFile("requests.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a JSON handler writing to the file (or use slog.NewTextHandler)
	handler := slog.NewJSONHandler(file, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	logger := slog.New(handler)

	// Optional: set as the global default logger
	slog.SetDefault(logger)
	cfg.logger = logger

	startrepl(&cfg)
}

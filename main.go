package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/embeque/pokedexcli/internal/pokecache"
)

const pokeapi = "https://pokeapi.co/api/v2/"

type config struct {
	registry  map[string]cliCommand
	Next      string
	Prev      string
	Cache     pokecache.Cache
	logger    *slog.Logger
	arguments []string
	personXp  int
}

func main() {
	// fmt.Println("Hello, World!")

	var cfg config
	cfg.registry = map[string]cliCommand{
		"map": {
			name:        "map",
			description: "List the Location Areas of Map",
			callback:    NextMapAction,
		},
		"mapb": {
			name:        "mapb",
			description: "List the previous Location Areas of Map",
			callback:    PrevMapAction,
		},
		"explore": {
			name:        "explore",
			description: "Explore the given Area for Pokemons",
			callback:    ExploreAction,
		},
		"catch": {
			name:        "catch",
			description: "Catch the given Pokemon",
			callback:    CatchAction,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect the given Pokemon",
			callback:    InspectAction,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all the cateched pokemons",
			callback:    PokedexAction,
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
	file, err := os.OpenFile("logs/requests.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
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

	// experience
	cfg.personXp = 0

	startrepl(&cfg)
}

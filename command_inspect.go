package main

import "fmt"

func print_pokemon_info(info pokemonInfo) {
	fmt.Printf("Name: %s\n", info.Name)
	fmt.Printf("Height: %d\n", info.Height)
	fmt.Printf("Weight: %d\n", info.Weight)
	fmt.Println("Stats:")
	for _, stat := range info.Stats {
		fmt.Printf("-%s: %d\n", stat.Stat.Name, stat.Value)
	}
	fmt.Println("Types:")

	for _, typ := range info.Types {
		fmt.Printf("-%s\n", typ.Type.Name)
	}
}

func InspectAction(cfg *config) error {
	pokemon := cfg.arguments[0]
	info, ok := Pokedex[pokemon]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	print_pokemon_info(info)
	return nil
}

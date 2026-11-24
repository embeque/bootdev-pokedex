package main

import "fmt"

func PokedexAction(cfg *config) error {
	fmt.Println("Your Pokedex:")
	for key, _ := range Pokedex {
		fmt.Printf("\t%s\n", key)
	}

	return nil
}

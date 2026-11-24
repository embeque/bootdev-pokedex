package main

import (
	"encoding/json"
	"fmt"
)

type area struct {
	EncounterAreas []struct {
		Pokemon struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func getPokemonsArea(data []byte) (area, error) {
	var area area
	err := json.Unmarshal(data, &area)
	if err != nil {
		return area, fmt.Errorf("DecodingError: %v\n", err)
	}
	return area, nil
}

func print_pokemons(area area) {
	for _, encounter := range area.EncounterAreas {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}
}

func ExploreAction(cfg *config) error {
	location_area := cfg.arguments[0]
	fmt.Printf("Exploring %s...\n", location_area)

	url := "https://pokeapi.co/api/v2/location-area/" + location_area
	var err error

	data, ok := cfg.Cache.Get(url)
	if !ok {
		cfg.logger.Info("URL Request: " + url)
		data, err = getBytes(url)
		cfg.Cache.Add(url, data)
		if err != nil {
			return err
		}
	} else {
		cfg.logger.Info("URL Cached : " + url)
	}
	area, err := getPokemonsArea(data)
	if err != nil {
		return err
	}

	print_pokemons(area)
	return nil
}

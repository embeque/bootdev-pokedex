package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

var Pokedex map[string]pokemonInfo = make(map[string]pokemonInfo)

type pokemonInfo struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		Value int `json:"base_stat"`
		Stat  struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

func getPokemonInfo(data []byte) (pokemonInfo, error) {
	var info pokemonInfo
	err := json.Unmarshal(data, &info)
	if err != nil {
		return info, fmt.Errorf("DecodingError: %v\n", err)
	}
	return info, nil
}

func catchPokemon(info pokemonInfo) {
	Pokedex[info.Name] = info
}

func CatchAction(cfg *config) error {
	pokemon := cfg.arguments[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)

	url := "https://pokeapi.co/api/v2/pokemon/" + pokemon
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

	pokemoninfo, err := getPokemonInfo(data)
	if err != nil {
		return err
	}
	pokemonXp := pokemoninfo.BaseExperience

	// if we are able to catch it
	randomXp := rand.Intn(pokemonXp)
	if cfg.personXp+randomXp > pokemonXp {
		catchPokemon(pokemoninfo)
		fmt.Printf("%s was caught!\n", pokemon)
	} else {
		fmt.Printf("%s escaped!\n", pokemon)
		cfg.personXp += 5
	}

	return nil
}

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type locationData struct {
	Count    int                 `json:"count"`
	Next     string              `json:"next"`
	Previous string              `json:"previous"`
	Results  []map[string]string `json:"results"`
}

func getBytes(url string) ([]byte, error) {
	var data []byte

	res, err := http.Get(url)
	if err != nil {
		return data, fmt.Errorf("RequestError: %v\n", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return data, fmt.Errorf("API returned status: %s\n", res.Status)
	}

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return data, fmt.Errorf("ResponseReadingError: %v\n", err)
	}

	return data, nil
}

func getLocation(data []byte) (locationData, error) {
	var location locationData
	err := json.Unmarshal(data, &location)
	if err != nil {
		return location, fmt.Errorf("DecodingError: %v\n", err)
	}

	// fmt.Printf("Got the data: %+v\n", location)

	return location, nil
}

func print_locations(locations []map[string]string) {
	for _, location := range locations {
		fmt.Println(location["name"])
	}
}

func NextMapAction(cfg *config) error {
	url := cfg.Next
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
	locations, err := getLocation(data)
	if err != nil {
		return err
	}

	if locations.Next != "" {
		cfg.Next = locations.Next
	} else { // next become null at the end of list
		cfg.Next = pokeapi + "location-area?offset=0&limit=20"
	}
	cfg.Prev = locations.Previous
	print_locations(locations.Results)

	return nil
}

func PrevMapAction(cfg *config) error {
	if cfg.Prev == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	url := cfg.Prev
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

	locations, err := getLocation(data)
	if err != nil {
		return err
	}

	cfg.Prev = locations.Previous
	// else {
	// 	cfg.Prev = pokeapi + "location-area?offset=0&limit=20"

	// }
	cfg.Next = locations.Next
	print_locations(locations.Results)

	return nil
}

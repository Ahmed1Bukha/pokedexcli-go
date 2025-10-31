package main

import (
	"encoding/json"
	"fmt"

	"github.com/ahmed1bukha/pokedexcli-go/poki_http"
)


func commandExplore(config *Config)error{
	fmt.Printf("Exploring %v...\n",config.Input)

	var result poki_http.LocationExploreResponse
	var err error
	cachedValue , hasValue := config.Cache.Get(config.url+config.Input)
	if hasValue{
	
		err = json.Unmarshal(cachedValue,&result)
		if err !=nil {
			return fmt.Errorf("couldn't unmarshal data in map command")
		}
	} else{
		result, err = poki_http.GetExploreArea(config.url,config.Input)
		if err !=nil {
			return fmt.Errorf("coudn't get from get http")
		}
	}

	fmt.Println("Found Pokemon:")
	for _,pokimon := range result.Encounters{
		fmt.Println("- ",pokimon.Pokemon.Name)
	}

	return nil
}
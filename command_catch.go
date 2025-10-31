package main

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/ahmed1bukha/pokedexcli-go/poki_http"
)

func commandCatch(config *Config)error{
	var result poki_http.PokemonResponse
	var err error
	cachedValue, hasValue := config.Cache.Get(config.Input)
	if hasValue {
		err = json.Unmarshal(cachedValue,&result)
		if err != nil{
			return fmt.Errorf("error marsheling json")
		}
	}else { 
		result,err = poki_http.GetPokemon(config.pokemonUrl,config.Input)
		if err !=nil {
			return fmt.Errorf("error getting the pokemon from http")
		}
		jsonData,err := json.Marshal(result)
		if err !=nil{
			return fmt.Errorf("coudn't get from get http")
		}
		config.Cache.Add(result.Name,jsonData)
	}
	fmt.Printf("Throwing a Pokeball at %v... \n",result.Name)
	catchChance := 100.0 / (100.0 + float64(result.Base_experience))
	randomRoll := rand.Float64()
	if randomRoll < catchChance {
		
		fmt.Printf("...%v was caught!\n", result.Name)
		config.CathedPokemons[result.Name]= result
	} else {
		
		fmt.Printf("...%v escaped!\n", result.Name)
	}
	return nil
}
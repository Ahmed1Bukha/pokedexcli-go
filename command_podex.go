package main

import "fmt"

func commandPodex(config *Config) error{
	podex:= config.CathedPokemons
	fmt.Println("Your Pokedex: ")
	for _,name := range podex{
		fmt.Printf("- %v \n",name.Name)
	}
	return nil
}
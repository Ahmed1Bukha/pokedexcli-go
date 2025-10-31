package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ahmed1bukha/pokedexcli-go/poki_cache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}
type Config struct{
	Next string
	Previous string
	url string
	Cache *poki_cache.Cache 
	Input string
}


func main(){
	cliCommands:= getCommands()
	cfg := &Config{
		Next: "https://pokeapi.co/api/v2/location-area/",
		Previous: "",
		url: "https://pokeapi.co/api/v2/location-area/",
		Input: "",
	 }
	 cfg.Cache = poki_cache.NewCache(time.Duration(time.Duration.Seconds(10)))
	
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words:= scanner.Text()
		if len(words) == 0 {
			continue
		}
		cleanedInput:= cleanInput(words)
		value,ok := cliCommands[cleanedInput[0]]
		if !ok{
			fmt.Println("Unknown command")
		} else{	
			if len(cleanedInput)>1 {
				cfg.Input = cleanedInput[1]
			}
				value.callback(cfg)
		
			
			
		}
	}
}

func getCommands()map[string]cliCommand{
	return map[string] cliCommand{
		"exit":{
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help":{
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map":{
			name:"map",
			description: "map the next 20 location",
			callback: commandMap,
		},
		"mapb":{
			name:"mapb",
			description: "map the previous 20 location",
			callback: commandMapb,
		},
		"explore":{
			name:"explore",
			description: "map the previous 20 location",
			callback: commandExplore,
		},
	}
}


func cleanInput(text string) []string{
	lowerdString := strings.ToLower(text)
	
	t := strings.TrimSpace(lowerdString)

	finalSlice:= strings.Split(t," ")
	return finalSlice
}




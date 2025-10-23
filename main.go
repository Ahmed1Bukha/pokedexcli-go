package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}


func main(){
	cliCommands:= getCommands()
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
			value.callback()
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
	}
}


func cleanInput(text string) []string{
	lowerdString := strings.ToLower(text)
	
	t := strings.TrimSpace(lowerdString)

	finalSlice:= strings.Split(t," ")
	return finalSlice
}




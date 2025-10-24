package main

import "fmt"


func commandHelp(config *Config)error{
	fmt.Println(`Welcome to the Pokedex!
Usage:
`)
for key,command:= range getCommands(){
	fmt.Println(key+": "+command.description)
}
	return nil
}
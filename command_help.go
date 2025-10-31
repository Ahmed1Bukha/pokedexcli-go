package main

import "fmt"


func commandHelp(config *Config)error{
	fmt.Printf(`Welcome to the Pokedex! \n Usage:\n`)
for key,command:= range getCommands(){
	fmt.Println(key+": "+command.description)
}
	return nil
}
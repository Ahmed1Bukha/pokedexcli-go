package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		word:= scanner.Text()
		cleanedInput:= cleanInput(word)
		fmt.Println("Your command was: "+cleanedInput[0])
	}
	
}

func cleanInput(text string) []string{
	lowerdString := strings.ToLower(text)
	
	t := strings.TrimSpace(lowerdString)

	finalSlice:= strings.Split(t," ")

	

	return finalSlice
}
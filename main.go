package main

import (
	"fmt"
	"strings"
)

func main(){
	output:= cleanInput("Hello World")

	fmt.Println(output)
	fmt.Println(len(output))
}

func cleanInput(text string) []string{
	lowerdString := strings.ToLower(text)
	
	t := strings.TrimSpace(lowerdString)

	finalSlice:= strings.Split(t," ")

	

	return finalSlice
}
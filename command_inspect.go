package main

import "fmt"

func commandInspect (config *Config)error{
	poki,ok:= config.CathedPokemons[config.Input]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return fmt.Errorf("you have not caught that pokemon")
	}
	fmt.Printf("Name: %v\n",poki.Name)
	fmt.Printf("Height: %v\n",poki.Height)
	fmt.Printf("Weight: %v\n",poki.Weight)
	fmt.Println("Stats: ")
	for _,stats := range poki.Stats{
		fmt.Printf("%v\n",stats)
	}



	return nil
}
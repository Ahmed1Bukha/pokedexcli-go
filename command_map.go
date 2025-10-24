package main

import (
	"fmt"

	"github.com/ahmed1bukha/pokedexcli-go/poki_http"
)

func commandMap(configMap *Config) error{
	result , err :=poki_http.GetLocationArea(configMap.Next)
	if err !=nil{
		return fmt.Errorf("error wiht getting map, %w",err)
	}
	configMap.Next = result.Next
	configMap.Previous = result.Previous
	for _,location := range result.Result{
		fmt.Println(location.Name)
	}
	
return nil
}
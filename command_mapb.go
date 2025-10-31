package main

import (
	"encoding/json"
	"fmt"

	"github.com/ahmed1bukha/pokedexcli-go/poki_http"
)

func commandMapb(configMap *Config) error{
	var result poki_http.LocationAreaResponse
	cachedValue,hasValue:= configMap.Cache.Get(configMap.Previous)
	var err error
	
	if hasValue{
	
		err = json.Unmarshal(cachedValue,&result)
		if err !=nil {
			return fmt.Errorf("couldn't unmarshal data in map command")
		}
	} else {
		result , err =poki_http.GetLocationArea(configMap.Previous)
		if err !=nil{
			return fmt.Errorf("error with getting map, %w",err)
		}
		jsonData,err := json.Marshal(result)
		if err != nil{
			return fmt.Errorf("error with marshal data %w",err)
		}
		configMap.Cache.Add(configMap.Previous,jsonData)
	}
	
	configMap.Next = result.Next
	configMap.Previous = result.Previous
	for _,location := range result.Result{
		fmt.Println(location.Name)
	}
	
return nil
}
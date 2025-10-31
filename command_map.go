package main

import (
	"fmt"

	"encoding/json"

	"github.com/ahmed1bukha/pokedexcli-go/poki_http"
)

func commandMap(configMap *Config) error{
	var result poki_http.LocationAreaResponse
	cachedValue,hasValue:= configMap.Cache.Get(configMap.Next)
	var err error
	
	if hasValue{
	
		err = json.Unmarshal(cachedValue,&result)
		if err !=nil {
			return fmt.Errorf("couldn't unmarshal data in map command")
		}
	} else {
		result , err =poki_http.GetLocationArea(configMap.Next)
		if err !=nil{
			return fmt.Errorf("error with getting map, %w",err)
		}
		jsonData,err := json.Marshal(result)
		if err != nil{
			return fmt.Errorf("error with marshal data %w",err)
		}
		configMap.Cache.Add(configMap.Next,jsonData)
	}
	
	configMap.Next = result.Next
	configMap.Previous = result.Previous
	for _,location := range result.Result{
		fmt.Println(location.Name)
	}
	
return nil
}

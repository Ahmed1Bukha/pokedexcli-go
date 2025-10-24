package poki_http

import "fmt"


type LocationAreaResponse struct{
	Count int `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Result [] LocationArea `json:"results"`
}

type LocationArea struct{
	Name string `json:"name"`
	Url string `json:"url"`
}

func GetLocationArea(url string) (LocationAreaResponse,error){
	locationAreas,err:= getHttp[LocationAreaResponse](url)
	if err !=nil {
		return LocationAreaResponse{}, fmt.Errorf("error with getting the area location")
	}
	
	return *locationAreas,nil
}
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

type LocationExploreResponse struct{
	Name string `json:"name"`
	Encounters []PokemonEncounter `json:"pokemon_encounters"`
}


type PokemonEncounter struct {
	Pokemon NamedResource `json:"pokemon"`
}
type NamedResource struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonResponse struct{
	Name string `json:"name"`
	Base_experience int `json:"base_experience"`
	Height int `json:"height"`
	Order int `json:"order"`
	Weight int `json:"weight"`
	Forms []PokemonFroms `json:"forms"`
	Moves []PokemonMoves`json:"moves"`
	Stats []PokemonStats `json:"stats"`

}
type PokemonFroms struct {
	Name string `json:"name"`
}
type PokemonMoves struct {
	Name string `json:"name"`
}
type PokemonStats struct {
	Base_stat int `json:"base_stat"`
	Effort int `json:"effort"`
	Stat Stat `json:"stat"`
}
type Stat struct{
	Name string `json:"name"`
}

func GetLocationArea(url string) (LocationAreaResponse,error){
	locationAreas,err:= getHttp[LocationAreaResponse](url)
	if err !=nil {
		return LocationAreaResponse{}, fmt.Errorf("error with getting the area location")
	}
	
	return *locationAreas,nil
}

func GetExploreArea(url string ,area string) (LocationExploreResponse, error){
	exploreArea,err := getHttp[LocationExploreResponse](url+area)
	if err !=nil {
		return LocationExploreResponse{},fmt.Errorf("error with getting explore location")
	}
	return *exploreArea,nil
}

func GetPokemon(url string ,name string) (PokemonResponse, error){
	exploreArea,err := getHttp[PokemonResponse](url+name)
	if err !=nil {
		return PokemonResponse{},fmt.Errorf("error with getting explore location")
	}
	return *exploreArea,nil
}
package poki_http

import (
	"encoding/json"
	"fmt"
	"net/http"
)




func getHttp [T any] (url string) (*T,error){
	resp,err := http.Get(url)
	if err !=nil{
		return nil, fmt.Errorf("Error with the request %v",err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
    }

	decoder:= json.NewDecoder(resp.Body)

	var result T

	err = decoder.Decode(&result)

	if err !=nil{
		return nil,fmt.Errorf("failed to decode to json: %w",err)
	}
	return &result,nil
}
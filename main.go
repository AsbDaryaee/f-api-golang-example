package main

import (
	"fmt"
	"httpcall/api"
)

func main() {

	randomFactResponse, randomFactError := api.GetARandomFact()

	if randomFactError != nil {
		fmt.Println(randomFactError)
	} else {
		fmt.Printf("--> Random Fact: %+v \n", *randomFactResponse)
	}

	factByIdResponse, factByIdError := api.GetFactById(16)

	if factByIdError != nil {
		fmt.Println(factByIdError)
	} else {
		fmt.Printf("--> Fact By ID: %+v \n", *factByIdResponse)
	}
}

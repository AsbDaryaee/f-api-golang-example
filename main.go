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

	FactByIdResponse, FactByIdError := api.GetFactById(16)

	if FactByIdError != nil {
		fmt.Println(FactByIdError)
	} else {
		fmt.Printf("--> Fact By ID: %+v \n", *FactByIdResponse)
	}
}

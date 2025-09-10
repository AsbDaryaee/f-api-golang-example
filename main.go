package main

import (
	"fmt"
	"httpcall/api"
)

func main() {

	fmt.Println("----> Fetching Data Started")
	defer fmt.Println("----> Fetching Data Finished")

	response, err := api.GetARandomFact()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v \n", *response)
	}
}

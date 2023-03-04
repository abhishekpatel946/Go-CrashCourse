package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps in go")

	language := make(map[string]string)

	language["js"] = "Javascript"
	language["go"] = "Golang"
	language["sol"] = "Solidity"

	fmt.Println("Mapping languages are:", language)
	fmt.Println("js shorts for :", language["js"])

	delete(language, "js")
	fmt.Println("Mapping languages are:", language)


	// loops 
	for key, value := range language {
		fmt.Printf("For key %v and value is %v \n", key, value)
	}

}

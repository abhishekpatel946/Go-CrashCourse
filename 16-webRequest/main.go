package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const URL = "https://lco.dev/"

func main() {
	fmt.Println("Welcome to web-request in go")

	res, err := http.Get(URL)
	checkError(err)

	fmt.Printf("Raw reponse type from http-get %T \n", res)
	fmt.Println("")
	fmt.Println("Raw reponse in byte from http-get \n", res)
	fmt.Println("")
	fmt.Println("response", res.StatusCode, res.Body)
	fmt.Println("")

	databytes, err := ioutil.ReadAll(res.Body)
	checkError(err)

	fmt.Println("databyte from response.body :", databytes)
	fmt.Println("")
	fmt.Println("databyte from response.body :", string(databytes))
	fmt.Println("")
	
	// caller's responsibility to close the connection
	defer res.Body.Close()

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

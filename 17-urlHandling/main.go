package main

import (
	"fmt"
	"net/url"
)

const MY_URL = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=qwerty123"

func main() {
	fmt.Println("Welcome to URL Handling in go")
	fmt.Printf("\n")

	// parsing
	result, err := url.Parse(MY_URL)
	checkError(err)

	fmt.Println("MY_URL : ", result)
	fmt.Printf("\n")
	fmt.Println("scheme : ", result.Scheme)
	fmt.Printf("\n")
	fmt.Println("host : ", result.Host)
	fmt.Printf("\n")
	fmt.Println("hostname : ", result.Hostname())
	fmt.Printf("\n")
	fmt.Println("port : ", result.Port())
	fmt.Printf("\n")
	fmt.Println("path : ", result.Path)
	fmt.Printf("\n")
	fmt.Println("query : ", result.RawQuery)
	fmt.Printf("\n")

	qparams := result.Query()
	fmt.Printf("the type of query params are: %T", qparams)
	fmt.Printf("\n")

	for _, value := range qparams {
		fmt.Println("value is: ", value)
	}
	fmt.Printf("\n")

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=abhishek",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println("anotherUrl is :", anotherUrl)
	fmt.Printf("\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

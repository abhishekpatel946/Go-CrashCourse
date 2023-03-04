package main

import "fmt"

func main() {
	fmt.Println("Welcome to defer in go")

	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")

	defer fmt.Println("World")
	fmt.Println("Hello")

	myDeferFunc()
}

func myDeferFunc() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}

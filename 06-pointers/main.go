package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	// var ptr *int
	// fmt.Println("Default value of pointer is:", ptr)	// <nil>

	myNumber := 23
	var ptr *int = &myNumber
	fmt.Println("Value of actual pointer is:", ptr)		// 0xc000016088
	fmt.Println("Value of actual pointing to:", *ptr)	// 23

	*ptr = *ptr * 2
	fmt.Println("new value is:", myNumber)	// 46

}

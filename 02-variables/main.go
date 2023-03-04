package main

import "fmt"


// Constant value with public accessible
const LoginToken string = "randomcrypto123456#$$%"
const privateVaraible = "something private"

func main() {

	// using var keyword with types
	var username string = "abhishek"
	fmt.Printf("\n\n The username is : %v", username)
	fmt.Printf("\n The type of username is : %T", username)

	var isLoggedIn bool = true
	fmt.Printf("\n\n The isLoggedIn is : %v", isLoggedIn)
	fmt.Printf("\n The type of isLoggedIn is : %T", isLoggedIn)

	var smallInt uint8 = 255
	fmt.Printf("\n\n The smallInt is : %v", smallInt)
	fmt.Printf("\n The type of smallInt is : %T", smallInt)

	var regularInt int = 1234567890111221357
	fmt.Printf("\n\n The regularInt is : %v", regularInt)
	fmt.Printf("\n The type of regularInt is : %T", regularInt)

	var smallFloat float32 = 255.45612379
	fmt.Printf("\n\n The smallFloat is : %v", smallFloat)
	fmt.Printf("\n The type of smallFloat is : %T", smallFloat)


	// /default values and some aliases
	var undefinedInt int
	fmt.Printf("\n\n The undefinedInt is : %v", undefinedInt)
	fmt.Printf("\n The type of undefinedInt is : %T", undefinedInt)

	var undefinedFloat float64
	fmt.Printf("\n\n The undefinedFloat is : %v", undefinedFloat)
	fmt.Printf("\n The type of undefinedFloat is : %T", undefinedFloat)

	var undefinedString string
	fmt.Printf("\n\n The undefinedString is : %v", undefinedString)
	fmt.Printf("\n The type of undefinedString is : %T", undefinedString)


	// implicit types
	var website = "google.com"
	fmt.Println("\n\n website ===", website)


	// no var style, using walras operator
	numberOfUsers := 40000
	fmt.Println("\n\n walras :=", numberOfUsers)


	// public global token
	fmt.Println("\n\n Public Global Variable ===", LoginToken)

	// public global token
	fmt.Println("\n\n Private Global Variable ===", privateVaraible)
}

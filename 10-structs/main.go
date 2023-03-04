package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs in go")
	// no inheritance in go; no super or parent-child

	user := User{"\n Abhishek Patel", "abhishekpatel946@gmail.com", true, 20}
	fmt.Println(user)

	fmt.Printf("\n User details are: %+v", user)

	fmt.Println(user.Name, user.Email)

}

type User struct {
	Name       string
	Email      string
	isVerified bool
	Age        int
}

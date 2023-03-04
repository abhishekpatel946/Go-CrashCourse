package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in go")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[2] = "Banana"
	fruitList[3] = "Cheery"

	fmt.Println("Value of array is:", fruitList)
	fmt.Println("accessing via indexes:", fruitList[0])
	fmt.Println("length of array", len(fruitList))

	var vegList = [3]string {"potato", "beans", "mushroom"}
	fmt.Println("value of vegList are:", vegList)

}

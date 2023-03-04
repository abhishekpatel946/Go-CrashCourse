package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in go")

	var fruitList = []string{"Apple", "Tomato"}

	fmt.Printf("Type of fruitList is %T", fruitList)
	fmt.Println("Values are:", fruitList)
	fmt.Println("length of fruitList:", len(fruitList))

	fruitList = append(fruitList, "Mango")
	fruitList = append(fruitList, "Banana", "Coconut")

	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	// "make" for raw memory management
	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 897
	highScores[2] = 654
	highScores[3] = 987
	// highScores[4] = 777

	fmt.Println(highScores)

	highScores = append(highScores, 777, 321, 555)

	fmt.Println(highScores)
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))


	// how to remove a value from slices based on index
	// [234 321 555 654 777 897 987] -> delete the 555 - 2nd index
	highScores = append(highScores[:2], highScores[2+1:]...)
	fmt.Println(highScores)

}

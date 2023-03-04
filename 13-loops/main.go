package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in go")

	days := []string{"sunday", "monday", "tuesday", "wednesday", "thurday", "friday", "saturday"}
	fmt.Println("The week days are: \n", days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("On index %v the value is %v", index, day)
	// 	fmt.Printf("\n")
	// }

	// for _, day := range days {
	// 	fmt.Println(day)
	// }

	rougueValue := 1

	for rougueValue < 10 {

		if rougueValue == 2 {
			rougueValue += 1
			continue
		}

		if rougueValue == 7 {
			goto githubsite
		}

		// if rougueValue == 5 {
		// 	break
		// }

		fmt.Println("value is :", rougueValue)
		rougueValue += 1
	}

githubsite:
	fmt.Println("Redirecting at github.com")

}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the rating for out Pizza:")

	// comma ok | err
	rating, _ := reader.ReadString('\n')
	// rating, err := reader.ReadString('\n')
	fmt.Println("Thanks for rating", rating)
	fmt.Printf("Type of this rating is %T", rating)
}

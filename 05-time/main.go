package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcom to time study of golang")

	currentTime := time.Now()
	fmt.Println("Curren time is:", currentTime)

	fmt.Println(currentTime.Format("01-02-2006 01:04:05 Monday"))

	createdDate := time.Date(2020, time.January, 10, 12, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
 
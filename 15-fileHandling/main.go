package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to file handling in go")

	content := "This needs to go in a file - github.com/abhishekpatel946"

	file, err := os.Create("./myfile.txt")
	checkErr(err)

	lenght, err := io.WriteString(file, content)
	checkErr(err)

	fmt.Println("lenght of file is: ", lenght)
	readFile("./myfile.txt")

	defer file.Close()
}

func readFile(filename string) {
	databytes, err := ioutil.ReadFile(filename)
	checkErr(err)

	fmt.Println("Byte data inside the file is: ", databytes)
	fmt.Println("Text data inside the file is: ", string(databytes))

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

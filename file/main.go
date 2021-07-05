package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("The File Practice!")

	content := "Hello from Go!"
	file, err := os.Create("./fromStr.txt")
	handleErrror(err)
	length, err := io.WriteString(file, content)
	handleErrror(err)
	fmt.Printf("Wrote the file with %v chars\n", length)
	defer file.Close()
	defer readFile("./fromStr.txt")

}

func readFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	handleErrror(err)
	fmt.Println("Text read from file:", string(data))
}

func handleErrror(err error) {
	if err != nil {
		panic(err)
	}

}

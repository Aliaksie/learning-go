package main

import (
	"fmt"
)

func main() {
	fmt.Println("The Arrays Prcatice!")

	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"

	fmt.Println(colors)
	fmt.Println(colors[1])

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	fmt.Println("Size of colores:", len(colors))
	fmt.Println("Size of numbers:", len(numbers))

}

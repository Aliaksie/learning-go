package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("The Slices Practice!")
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	colors = append(colors, "White")
	fmt.Println(colors)

	colors = append(colors[0 : len(colors)-1])
	fmt.Println(colors)

	numbers := make([]int, 5)
	fmt.Println(numbers)
	numbers[0] = 5
	numbers[1] = 4
	numbers[2] = 3
	numbers[3] = 2
	numbers[4] = 1
	fmt.Println(numbers)

	numbers = append(numbers, 6)
	fmt.Println(numbers)

	sort.Ints(numbers)
	fmt.Println(numbers)
}

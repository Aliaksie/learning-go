package main

import (
	"fmt"
)

func main() {
	fmt.Println("The Simple Functions Practice!")
	doSomething()
	sum := addValue(1, 2)
	fmt.Println("Add vaule ", sum)

	sum = addAllValue(1, 2, 3)
	fmt.Println("Add All vaule ", sum)

	var len int
	sum, len = addAllValueAndLen(1, 2, 3)

	fmt.Printf("Sum=%v, Len=%v", sum, len)

}

func doSomething() {
	fmt.Println("Doing")
}

func addValue(val1, val2 int) int {
	return val1 + val2
}

func addAllValue(values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}

func addAllValueAndLen(values ...int) (int, int) {
	total := 0
	for _, v := range values {
		total += v
	}
	return total, len(values)
}

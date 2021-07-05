package main

import (
	"fmt"
)

func main() {
	fmt.Println("The Pointers Practice!")

	var p *int
	fmt.Println("Value of p:", p)
	// fmt.Println("Value of p:", *p) crash nil pointer

	anInt := 42
	var p1 = &anInt
	fmt.Println("Value of p1:", *p1)

	val1 := 42.13
	p2 := &val1
	fmt.Println("Value of p2:", *p2)

	*p2 = *p2 / 31
	fmt.Println("Value of pointer:", *p2)
	fmt.Println("Value of original:", val1)
}

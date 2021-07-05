package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("The Conditional Practice")

	exampleIf()
	exampleSwitch()
	exampleLoops()
}

func exampleIf() {
	fmt.Println("The IF example")

	theAnswer := 42
	var rs string
	if theAnswer < 0 {
		rs = "Less than zero"
	} else if theAnswer == 0 {
		rs = "Equal to zero"
	} else {
		rs = "Greater than zero"
	}

	fmt.Println("Result ", rs)

	if x := -43; x < 0 {
		rs = "Less than zero"
	} else if x == 0 {
		rs = "Equal to zero"
	} else {
		rs = "Greater than zero"
	}
	fmt.Println("Rersults ", rs)
	fmt.Println("End IF example")
}

func exampleSwitch() {
	fmt.Println("The SWITCH example")
	rand.Seed(time.Now().Unix())
	dow := rand.Intn(7) + 1
	fmt.Println("Day", dow)
	// dow = 3
	var rs string
	switch dow {
	case 1:
		rs = "It's Sunday!"
	case 2:
		rs = "It's Monday!"
	case 3:
		rs = "It's Tuesday!"
		fmt.Println("Case:3 ", rs)
		fallthrough
	default:
		rs = "It's other day!"

	}
	fmt.Println("Results: ", rs)

	fmt.Println("End SWITCH example")
}

func exampleLoops() {
	fmt.Println("The Loops example")
	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	for i := range colors {
		fmt.Println(colors[i])
	}

	for _, color := range colors {
		fmt.Println(color)
	}

	value := 1
	for value < 10 {
		fmt.Println("Value=", value)
		value++
	}

	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Println("Sum:", sum)
		if sum > 200 {
			goto theEnd
		}

	}

theEnd:
	fmt.Println("End of Program")

	fmt.Println("End loops example")
}

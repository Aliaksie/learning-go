package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("An Advance Calculator!")

	reader := bufio.NewReader(os.Stdin)

	val1 := geValue(reader)
	val2 := geValue(reader)

	var rs float64
	switch operation := getOperations(reader); operation {
	case "+":
		rs = sum(val1, val2)
	case "-":
		rs = subtract(val1, val2)
	case "*":
		rs = multiply(val1, val2)
	case "/":
		rs = divide(val1, val2)
	default:
		panic("Invalid operation!")

	}

	rs = math.Round(rs*100) / 100
	fmt.Printf("The results is %v\n\n", rs)
}

func geValue(reader *bufio.Reader) float64 {
	fmt.Println("Please insert value")
	input, _ := reader.ReadString('\n')
	float, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
		panic("The input value must be number")
	}

	return float
}

func getOperations(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func sum(val1, val2 float64) float64 {
	return val1 + val2
}

func subtract(val1, val2 float64) float64 {
	return val1 - val2
}

func multiply(val1, val2 float64) float64 {
	return val1 * val2
}

func divide(val1, val2 float64) float64 {
	return val1 / val2
}

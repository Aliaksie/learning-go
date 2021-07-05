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
	fmt.Println("A Simple Sum Calculator!")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please insert first value")
	float1 := getFloat(reader)
	fmt.Println("Please insert second value")
	float2 := getFloat(reader)

	sum := float1 + float2
	sum = math.Round(sum*100) / 100
	fmt.Printf("The sum of %v and %v is %v\n\n", float1, float2, sum)
}

func getFloat(reader *bufio.Reader) float64 {
	input, _ := reader.ReadString('\n')
	float, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
		panic("The input value must be number")
	}

	return float
}

package main

import (
	"fmt"
)

// Dog is a struct
type Dog struct {
	Bread  string
	Weight int
}

type cat struct {
	Bread  string
	Weight int
}

func main() {
	fmt.Println("The Structs Practice!")

	poodle := Dog{"Poodle", 10}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	poodle.Weight = 8
	fmt.Printf("%+v\n", poodle)

	scottishFold := cat{"scottish fold", 10}
	fmt.Println(scottishFold)

}

package main

import (
	"fmt"
)

// Dog is struct
type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

// Speak is how the dog speaks
func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

// SpeakThreeTimes is how the dog speaks loundly
func (d Dog) SpeakThreeTimes() {
	d.Sound = fmt.Sprintf("%v %v %v", d.Sound, d.Sound, d.Sound)
	fmt.Println(d.Sound)
}

func main() {
	fmt.Println("The Custom Functions Practice!")

	poodle := Dog{"Poodle", 10, "Woof!"}
	fmt.Printf("%+v\n", poodle)
	poodle.Speak()
	poodle.Sound = "Arf!"
	poodle.Speak()
	poodle.SpeakThreeTimes()
	poodle.SpeakThreeTimes()
}

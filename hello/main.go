package main

import (
	"fmt"
)

const aConst int = 64

func main() {
	var aString string = "Hello from GO!"
	fmt.Println(aString)
	fmt.Printf("The var's  type is %T\n", aString)

	var anInteger int = 42
	fmt.Println(anInteger)
	fmt.Printf("The var's  type is %T\n", anInteger)

	var defaultInt int
	fmt.Println(defaultInt)
	fmt.Printf("The var's  type is %T\n", defaultInt)

	var anotherString = "This is another"
	fmt.Println(anotherString)
	fmt.Printf("The var's  type is %T\n", anotherString)

	var anotherInt = 53
	fmt.Println(anotherInt)
	fmt.Printf("The var's  type is %T\n", anotherInt)

	myString := "My string :="
	fmt.Println(myString)
	fmt.Printf("The var's  type is %T\n", myString)

	fmt.Println(aConst)
	fmt.Printf("The var's  type is %T\n", aConst)

}

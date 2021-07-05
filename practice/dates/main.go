package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Dates Practice!")

	t := time.Now()
	fmt.Println("The time is now at ", t)

	d := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("The Go launched at ", d)
	fmt.Println("The formated time ", d.Format(time.ANSIC))

	parsedTime, _ := time.Parse(time.ANSIC, "Tue Nov 10 23:00:00 2009")
	fmt.Println("The parsed time ", parsedTime)
}

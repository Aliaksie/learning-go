package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const url = "http://services.explorecalifornia.org/json/tours.php"

type Tour struct {
	Name, Price string
}

func main() {
	fmt.Println("The Web RQ Practice!")

	rs, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("RS type: %T\n", rs)
	defer rs.Body.Close()

	bytes, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		panic(err)
	}

	content := string(bytes)
	fmt.Println(content)

	tours := toursFromJson(content)
	for _, v := range tours {
		fmt.Printf("%+v\n", v)
	}
}

func toursFromJson(content string) []Tour {
	tours := make([]Tour, 0, 20)

	decoder := json.NewDecoder(strings.NewReader(content))

	_, err := decoder.Token()
	if err != nil {
		panic(err)
	}

	var tour Tour
	for decoder.More() {
		err := decoder.Decode(&tour)
		if err != nil {
			panic(err)
		}

		tours = append(tours, tour)
	}
	return tours
}

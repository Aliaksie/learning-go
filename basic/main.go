package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strings"
)

func fizzBuzz(number int) {
	fmt.Println("The number is", number)
	rs := ""

	if number%3 == 0 && number%5 == 0 {
		rs = "fizz buzz"
	} else if number%3 == 0 {
		rs = "fizz"
	} else if number%5 == 0 {
		rs = "buzz"
	}

	fmt.Println(rs)
}

func evenEndedNumber(start int, end int) {
	count := 0
	for a := start; a <= end; a++ {
		for b := a; b <= end; b++ {
			n := a * b
			s := fmt.Sprintf("%d", n)
			if s[0] == s[len(s)-1] {
				count++
			}

		}
	}
	fmt.Println("The count EvenEndedNumber is", count)
}

func findMax(nums []int) {
	max := nums[0]
	for _, num := range nums[1:] {
		if num > max {
			max = num
		}
	}
	fmt.Println("Tha max numbers is", max)
}

func worldCount(text string) {
	maps := map[string]int{}
	words := strings.Fields(text)
	for _, world := range words {
		maps[strings.ToLower(world)]++
	}
	fmt.Println("The worlds count are", maps)
}

func contentType(url string) (string, error) {
	rs, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer rs.Body.Close()

	ctype := rs.Header.Get("Content-Type")
	if ctype == "" {
		return "", fmt.Errorf("can't find Content-Type header")
	}

	return ctype, nil
}

type Point struct {
	X int
	Y int
}

func (p *Point) Move(dx int, dy int) {
	p.X += dx
	p.Y += dy
}

type Square struct {
	Center Point
	Length int
}

func NewSquare(x int, y int, length int) (*Square, error) {
	if length < 0 {
		return nil, fmt.Errorf("the `length` must be > 0")
	}

	s := &Square{
		Center: Point{x, y},
		Length: length,
	}

	return s, nil
}

func (s *Square) Move(dx int, dy int) {
	s.Center.Move(dx, dy)
}

func (s *Square) Area() float64 {
	return float64(s.Length) * float64(s.Length)
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

type Shape interface {
	Area() float64
}

type Capper struct {
	wrt io.Writer
}

func (c *Capper) Write(p []byte) (n int, err error) {
	diff := byte('a' - 'A')

	out := make([]byte, len(p))
	for i, b := range p {
		if b >= 'a' && b <= 'z' {
			b -= diff
		}
		out[i] = b
	}
	return c.wrt.Write(out)
}

func setUpLog() {
	out, err := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	defer out.Close()

	log.SetOutput(out)
}

type Config struct {
}

func readConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	cfg := &Config{}

	return cfg, nil
}

func main() {
	fmt.Println("Go Essential!")
	setUpLog()

	fizzBuzz(15)
	evenEndedNumber(10, 100)
	findMax([]int{10, 15, 5, 1, 30})

	text := `
	Needles and pins
	needles And Pins
	Sew me a sail
	To catch me the wind
	`
	worldCount(text)

	ctype, err := contentType("https://translate.google.com/")
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Println("The Content-Type is", ctype)
	}

	s, err := NewSquare(1, 1, 10)
	if err != nil {
		log.Fatalf("Error can't create square")
	}
	s.Move(2, 3)
	fmt.Printf("%+v\n", s)
	fmt.Println("The square area is ", s.Area())

	c := &Circle{3}
	fmt.Println("The circle area is ", c.Area())

	total := 0.0
	for _, s := range []Shape{s, c} {
		total += s.Area()
	}

	fmt.Println("The sum area is ", total)

	cfg, err := readConfig("config.txt")
	if err != nil {
		fmt.Printf("error: %s\n", err)
		log.Printf("error %+v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%T", cfg)

}

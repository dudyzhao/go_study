package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	num := rand.Intn(1000)
	fmt.Println("num == ", num)
	switch {
	case num == 0:
		fmt.Println("Fizz")
	case num%5 == 0:
		fmt.Println("Buzz")
	case num%3 == 0 && num%5 == 0:
		fmt.Println("FizzBuzz")
	default:
		fmt.Println(num)
	}
}

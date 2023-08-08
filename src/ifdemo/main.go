package main

import (
	"fmt"
)

func main() {
	x := 27
	if x%2 == 0 {
		fmt.Println(x, "is even")
	} else {
		fmt.Println(x, "not even")
	}

	if num := givenumber(); num < 0 {
		fmt.Println(num, "小于0")
	} else if num < 10 {
		fmt.Println(num, "小于10")
	} else {
		fmt.Println(num, "大于0")
	}
}

func givenumber() int {
	return -1
}

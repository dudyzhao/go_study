package main

import (
	"fmt"
)

func highlow(high int, low int) {
	if high < low {
		fmt.Println("panic")
		panic("highlow() low greater than high")
	}
	defer fmt.Println("defer:highlow(", high, ",", low, ")")
	fmt.Println("Call:hithlow(", high, ",", low, ")")
	highlow(high, low+1)
}

func main() {
	defer func() {
		handler := recover()
		if handler != nil {
			fmt.Println("main():recover", handler)
		}
	}()
	highlow(2, 0)
	fmt.Println("Program finished successed")
}

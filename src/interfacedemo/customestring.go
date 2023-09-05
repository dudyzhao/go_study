package main

import (
	"fmt"
)

type Person struct {
	Name, Country string
}

// 重写string的String方法
func (p Person) String() string {
	return fmt.Sprintf("%v is from %v", p.Name, p.Country)
}

func main() {
	us := Person{"john", "USA"}
	ja := Person{"adf", "Jap"}
	ch := Person{"张三", "China"}
	fmt.Printf("%s\n%s\n%s", us, ja, ch)
}

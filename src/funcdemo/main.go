package main

import (
	"fmt"
)

type triangle struct {
	size int
}
type square struct {
	size int
}

func (t triangle) perimeter() int {
	return t.size * 3
}

func (s square) perimeter() int {
	return s.size * 4
}

// 使用指针，可以修改struct的变量值
func (t *triangle) doublesize() {
	t.size = t.size * 2
}

func main() {
	t := triangle{3}
	s := square{4}
	fmt.Println("permeter (triangle):", t.perimeter())
	fmt.Println("permeter (square):", s.perimeter())
	t.doublesize()
	fmt.Println("t.size:", t.size)
	fmt.Println("permeter (triangle):", t.perimeter())
}

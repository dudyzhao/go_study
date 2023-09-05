package main

import (
	"fmt"
	"math"
)

// 定义接口
type Shape interface {
	Perimeter() float64
	Area() float64
}

// 实现接口
type Square struct {
	size float64
}

func (s Square) Area() float64 {
	return s.size * s.size
}
func (s Square) Perimeter() float64 {
	return s.size * 4
}

// 实现接口
type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func printInfomation(s Shape) {
	fmt.Printf("%T\n", s)
	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
}

func main() {
	s := Square{3}
	printInfomation(s)
	c := Circle{4}
	printInfomation(c)
}

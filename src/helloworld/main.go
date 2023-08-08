package main

import "fmt"

// import "strconv"

var firstName string

var (
	lastName string = "赵"
	age             = 20
)

// const HttpStatus = 200
// const (
// 	StatusOK    = 0
// 	StatusError = 2
// )

// func sum(number1 string ,number2,string) (result int) {
// 	int1,_ := strconv.Atoi(number1)
// 	int2,_ := strconv.Atoi(number2)
// 	result = int1 + int2
// 	return
// }

// func calc(number1 string,number2 string) (sum int, mul int){
// 	int1,_ := strconv.Atoi(number1)
// 	int2,_ := strconv.Atoi(number2)
// 	sum = int1 + int2
// 	mul = int1 * int2
// 	return
// }

var name = "dudy"

func updateName(name *string) {
	*name = "David"
}

func main() {
	fmt.Println("Hello world!")
	lastName := "dudy, "
	age := 30
	fmt.Println(firstName, lastName, age)
	updateName(&name)
	fmt.Println(name)
}

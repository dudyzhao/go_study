package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Println("Hey ~I'm a log message!")
	employee, error := getInformation(1001)
	if error != nil {
		fmt.Println("something is wrong")
	} else {
		fmt.Println(employee)
	}
	emaployee1, error := getInformation1(1000)
	if error != nil {
		fmt.Println("something is wrong")
	} else {
		fmt.Println(emaployee1)
	}

	emaployee2, error := getInformation2(1002)
	if errors.Is(error, ErrNotFound) {
		fmt.Println("NOT FOUND：%v\n", error)
	} else {
		fmt.Println(emaployee2)
	}
}

func getInformation(id int) (*Employee, error) {
	employee, error := apiCallEmployee(id)
	if error != nil {
		return nil, error //实体返回空，并将异常一同放回
		// return nil, fmt.Errorf("got an error when getting the emaployee infomation: %v",error) 或者返回一个加强版的异常
	}
	return employee, error
}

var ErrNotFound = errors.New("Employee not found!")

func getInformation2(id int) (*Employee, error) {
	employee, error := apiCallEmployee(id)
	if error == nil {
		return employee, error
	}
	return nil, ErrNotFound
}

func getInformation1(id int) (*Employee, error) {
	for tries := 0; tries < 3; tries++ {
		employee, error := apiCallEmployee(id)
		if error == nil {
			return employee, nil
		}
		fmt.Println("server is not responding,retrying...")
		time.Sleep(time.Second * 2)
	}
	return nil, fmt.Errorf("server has failed to respond to get the employee information")
}

func apiCallEmployee(id int) (*Employee, error) {
	employee := Employee{ID: id, LastName: "Deo", FirstName: "john"}
	return &employee, nil
}

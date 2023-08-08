package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += 1
	}
	fmt.Println(sum)

	var num int64
	rand.Seed(time.Now().Unix())
	for num != 5 { // 只要！=5，就一直循环，类似于while
		num = rand.Int63n(15)
		fmt.Println(num, "==5")
	}
	fmt.Println(num, "while")

	for num := 7; num != 1; num-- {
		if num == 2 {
			fmt.Println(num, "==2")
			break // 跳出循环
		}
		fmt.Println(num, "break")
	}

	for num := 7; num != 0; num-- {
		if num == 2 {
			fmt.Println(num, "continue ==2")
			continue // 继续下一轮循环
		}
		fmt.Println(num, "continue")
	}
}

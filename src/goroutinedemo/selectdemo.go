package main

import (
	"fmt"
	"time"
)

func process(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "Done process"
}

func replicate(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "Done replicate"
}

func main() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	go process(ch1)
	go replicate(ch2)
	for i := 0; i < 2; i++ {
		select { // 多路复用，类似于switch，需要搭配循环，否则当某个条件达成，则程序结束
		case process := <-ch1:
			fmt.Println(process)
		case replicate := <-ch2:
			fmt.Println(replicate)
		}
	}
}

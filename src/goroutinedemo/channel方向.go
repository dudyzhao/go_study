package main

import "fmt"

// chan<- string 只发送数据
func send(ch chan<- string, message string) {
	fmt.Println("我要发送数据", message)
	ch <- message
}

// <-chan string 只接受数据
func receive(ch <-chan string) {
	fmt.Println("我接受到的数据为：", <-ch)
}
func main() {
	ch := make(chan string, 1)
	send(ch, "我是数据")
	receive(ch)
}

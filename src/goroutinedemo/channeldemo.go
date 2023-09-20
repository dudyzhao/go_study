package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()
	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}
	// 指定channel队列大小，目前测试，channel队列越大，效率越高
	ch := make(chan string, 1000)
	for _, api := range apis {
		// 每次都创建一个新的goroutine
		go checkApi(api, ch)
		fmt.Printf(<-ch)
	}
	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}

func checkApi(api string, ch chan string) {
	_, err := http.Get(api)
	if err != nil {
		ch <- fmt.Sprintf("ERROR: %s is down!\n", api)
		return
	}
	ch <- fmt.Sprintf("SUCCESS: %s is up and running!\n", api)
}

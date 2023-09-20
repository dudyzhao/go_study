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
	//for _, api := range apis {
	//	_, err := http.Get(api)
	//	if err != nil {
	//		fmt.Printf("Error: %s is down !\n", api)
	//		continue
	//	}
	//	fmt.Printf("Sucess: %s is up and running !\n", api)
	//}
	for _, api := range apis {
		go checkApi(api)
	}
	elapsed := time.Since(start)
	time.Sleep(10 * time.Second)
	fmt.Printf("Done ! It look %v seconds ! \n", elapsed.Seconds())
}

func checkApi(api string) {
	_, err := http.Get(api)
	if err != nil {
		fmt.Printf("Error: %s is down !\n", api)
		return
	}
	fmt.Printf("Sucess: %s is up and running !\n", api)
}

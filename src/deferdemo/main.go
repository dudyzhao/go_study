package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	for i := 1; i <= 4; i++ {
		defer fmt.Println("defer", -i)
		fmt.Println("result", i)
	}

	newfile, error := os.Create("learnGo.txt")
	if error != nil {
		fmt.Println("error:could not create file")
		return
	}
	defer newfile.Close()

	if _, error = io.WriteString(newfile, "Learn Go!"); error != nil {
		fmt.Println("Error: Could not write to file")
		return
	}
	newfile.Sync()
	fmt.Println("this is over")
}

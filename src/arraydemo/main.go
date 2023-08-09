package main

import (
	"fmt"
)

func main() {
	var a [3]int
	a[1] = 10
	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[len(a)-1])
	city := [5]string{"one", "tow", "three", "four"} // 指定数组大小，且指定数据集，当数据集不足，补空格
	fmt.Println("cifies", city)
	q := [...]string{"New York", "Paris", "Berlin", "Madrid"} // 不指定数组大小，但必须指定数据集
	fmt.Println("q = ", q)
	numbers := [...]int{99: -1} // 不指定数组大小，也不指定具体数据集，但指定最后一个数据集的值
	fmt.Println("numbers = ", numbers)

	var twoD [3][5]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			twoD[i][j] = (i + 1) * (j + 1)
		}
		fmt.Println("Row", i, twoD[i])
	}
	fmt.Println("all in one ", twoD)
}

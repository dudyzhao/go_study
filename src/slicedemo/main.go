package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 切片定义可以和数组一模一样
	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	fmt.Println("slice len : ", len(months), "slice cap : ", cap(months), months)
	// 切片基本写法
	// 数组[i:p]：i--切片第一个元素，对应months[i]的元素; p--要使用数组的元素数目
	quarter2 := months[3:6]
	// 切片长度固定--p，但容量是从i -- 数组的最后一个元素下标为止，只是不可见。容量是指切片可以扩容的程度
	fmt.Println("quarter2:", quarter2)
	for i := 0; i < 12; i++ {
		quarter2 = append(quarter2, strconv.Itoa(i))
	}
	fmt.Println("quarter2", quarter2)

	remove := 2
	if remove < len(months) {
		fmt.Println("Before:",months,"Remove ",months[remove])
		months = append(months[:remove],months[remove+1:]...)
		fmt.Println("after ",months)
	}

	slice2 := make([]string,3)
	copy(slice2,months)
	fmt.Println(slice2)
}

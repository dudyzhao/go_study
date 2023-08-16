package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")
	// 声明和初始化map
	studentsAge := map[string]int{
		"join": 32,
		"bob":  31,
	}
	fmt.Println(studentsAge)
	// 推荐的声明map
	studentsAge1 := make(map[string]int)
	// 给map添加元素
	studentsAge1["join"] = 32
	fmt.Println(studentsAge1)
	// 获取map中项的值
	fmt.Println("join's age is ", studentsAge1["join"])
	// 获取map中不存在的项的值时，返回默认值
	fmt.Println("bob's age is ", studentsAge1["bob"])
	// 判断map中项是否存在
	age, exist := studentsAge1["bob"]
	if exist {
		fmt.Println("bob's age is ", age)
	} else {
		fmt.Println("bob's age is not found")
	}
	// 删除项
	delete(studentsAge, "bob")
	fmt.Println("studentsAge is ", studentsAge)
	// map循环，先生成项，然后生成该项的值。不想输出项，可以使用_，如for _,age := range studentsAge{}
	studentsAge["bob"] = 31
	for name, age := range studentsAge {
		fmt.Printf("%s\t%v\n", name, age)
	}
	// 只循环map中的项
	for name := range studentsAge {
		fmt.Printf("name is %v", name)
	}
}

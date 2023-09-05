package main

import (
	"fmt"
)

func main() {
	x := 27
	if x%2 == 0 {
		fmt.Println(x, "is even")
	} else {
		fmt.Println(x, "not even")
	}

	if num := givenumber(); num < 0 {
		fmt.Println(num, "小于0")
	} else if num < 10 {
		fmt.Println(num, "小于10")
	} else {
		fmt.Println(num, "大于0")
	}
	fmt.Println("MCMZ is", romanToArabic("MCMZ"))
}

func givenumber() int {
	return -1
}

func romanToArabic(numeral string) int {
	romanMap := map[rune]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}

	arabicVals := make([]int, len(numeral)+1)

	for index, digit := range numeral {
		if val, present := romanMap[digit]; present {
			arabicVals[index] = val
		} else {
			fmt.Printf("Error: The roman numeral %s has a bad digit: %c\n", numeral, digit)
			return 0
		}
	}
	fmt.Println("arabicVals1111 is ", arabicVals)
	total := 0

	for index := 0; index < len(numeral); index++ {
		if arabicVals[index] < arabicVals[index+1] {
			arabicVals[index] = -arabicVals[index]
		}
		total += arabicVals[index]
	}

	return total
}

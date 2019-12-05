package main

import "fmt"

//练习5.15：编写类似sum的可变参数函数max和min。考虑不传参时，max和min该如何处理，再编写至少接收1个参数的版本。

func max(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	m := nums[0]
	for _, num := range nums {
		if m < num {
			m = num
		}
	}
	return m
}

func min(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	m := nums[0]
	for _, num := range nums {
		if m > num {
			m = num
		}
	}
	return 0
}

func max2(first int, nums ...int) int {
	m := first
	for _, num := range nums {
		if m < num {
			m = num
		}
	}
	return m
}

func min2(first int, nums ...int) int {
	m := first
	for _, num := range nums {
		if m > num {
			m = num
		}
	}
	return 0
}

func main() {
	fmt.Println(min(3, -1, 4))
	fmt.Println(max(3, -1, 4))
	fmt.Println(min2(3, -1, 4))
	fmt.Println(max2(3, -1, 4))
}

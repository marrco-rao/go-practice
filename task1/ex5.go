package main

import "fmt"

/*
加一：
给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。
将大整数加 1，并返回结果的数字数组。
*/
// 参考
func plusOne1(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	digits = append([]int{1}, digits...)
	return digits
}

func plusOne(digits []int) []int {
	// 输入空数组
	if len(digits) == 0 {
		return append(digits, 1)
	}
	// 条件 0 <= digits[i] <= 9 暂时不在题目里做判断
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i]++
			return digits
		}
	}
	//	循环完后，向前进行补位
	digits = append([]int{1}, digits...)
	return digits
}

func ex5_test() {
	d1 := []int{1, 2, 3}
	d2 := []int{4, 3, 2, 1}
	d3 := []int{9}
	d4 := []int{2, 3, 9}
	d5 := []int{9, 9, 9, 9}
	d6 := []int{}
	d7 := make([]int, 0)
	fmt.Println(plusOne(d1))
	fmt.Println(plusOne(d2))
	fmt.Println(plusOne(d3))
	fmt.Println(plusOne(d4))
	fmt.Println(plusOne(d5))
	fmt.Println(plusOne(d6))
	fmt.Println(plusOne(d7))
}

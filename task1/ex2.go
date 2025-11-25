package main

import (
	"fmt"
	"math"
	"strconv"
)

/*
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
例如，121 是回文，而 123 不是。
*/
func isPalindrome(x int) bool {
	// 排除负数
	if x < 0 {
		return false
	}
	str := strconv.Itoa(x) // 先转为字符串 "12345"
	runes := []rune(str)
	len_x := len(runes)
	fmt.Println(len_x / 2)
	// len_x/2 会自动丢弃小数，len_x/2 就是回数的前半段
	for i := 0; i < len_x/2; i++ {
		if runes[i] != runes[len_x-1-i] {
			return false
		}
	}
	return true
}

func ex2_test1() {
	x1 := 1001
	x2 := 12321
	x3 := 10
	x4 := -10
	fmt.Println("x1: ", isPalindrome(x1))
	fmt.Println("x2: ", isPalindrome(x2))
	fmt.Println("x3: ", isPalindrome(x3))
	fmt.Println("x4: ", isPalindrome(x4))
	fmt.Println("End ex2_test1")
}

// 尝试不用转字符串来处理
func isPalindrome2(x int) bool {
	// 排除负数
	if x < 0 {
		return false
	}
	// 找出整数长度
	len_x := IntLength(x)
	// fmt.Println(1234%10, "-", (1234/10)%10, "-", (1234/100)%10, "-", (1234/1000)%10)
	for i := 0; i <= len_x/2; i++ {
		//fmt.Println(i, " 位置 ", len_x-1-i)
		low_num := (x / int(math.Pow10(i))) % 10
		high_num := (x / int(math.Pow10(len_x-1-i))) % 10
		if low_num != high_num {
			fmt.Println(low_num, high_num)
			return false
		}
	}
	return true
}

// 循环法计算int长度
func IntLength(n int) int {
	if n == 0 {
		return 1
	}
	count := 0
	for n != 0 {
		n = n / 10
		count++
	}
	return count
}

// 对数法取int长度: AI分析推荐，速度快
func IntLength1(n int) int {
	if n == 0 {
		return 1
	}
	if n < 0 {
		n = -n // 处理负数
	}
	return int(math.Log10(float64(n))) + 1
}

func ex2_test2() {
	x1 := 1001
	x2 := 12321
	x3 := 10
	x4 := -10
	fmt.Println("x1: ", isPalindrome2(x1))
	fmt.Println("x2: ", isPalindrome2(x2))
	fmt.Println("x3: ", isPalindrome2(x3))
	fmt.Println("x4: ", isPalindrome2(x4))
	fmt.Println("End ex2_test2")
}

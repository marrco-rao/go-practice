package main

import (
	"errors"
	"fmt"
)

/*
136. 只出现一次的数字
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
*/
func singleNumber(nums []int) (int, error) {
	fmt.Println("nums: ", nums) // nums:  []
	fmt.Println(nums == nil)    // true
	fmt.Println(len(nums))      // 0
	// fmt.Println(nums[0])        // panic

	// 参数空值处理一下，避免后面直接使用nums[0]这种导致Panic。
	// 不处理对本题无影响
	if nums == nil {
		return 0, errors.New("Parameter is nil")
	}

	numMap := make(map[int]int)
	// 统计每个元素出现的次数
	for _, v := range nums {
		numMap[v] = numMap[v] + 1
	}
	for num, count := range numMap {
		if count == 1 {
			return num, nil
		}
	}
	// 如果没有找到出现次数为1的元素，返回错误
	return 0, errors.New("no single number")
}

func ex1_test1() {
	num := []int{4, 1, 2, 1, 2}
	res, err := singleNumber(num)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
	fmt.Println("End ex1_test1!")
}

func ex1_test2() {
	num := []int{}
	res, err := singleNumber(num)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
	fmt.Println("End ex1_test2!")
}
func ex1_test3() {
	res, err := singleNumber(nil)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
	fmt.Println("End ex1_test3!")
}

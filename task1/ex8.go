package main

import "fmt"

/*
基础
两数之和
考察：数组遍历、map使用
题目：给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数，并返回它们的数组下标。

示例 1：

输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
示例 2：

输入：nums = [3,2,4], target = 6
输出：[1,2]
示例 3：

输入：nums = [3,3], target = 6
输出：[0,1]
*/

// 常规遍历
func twoSum1(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 参考优化
func twoSum(nums []int, target int) []int {
	hash := map[int]int{}
	for i := 0; i < len(nums); i++ {
		x := nums[i]
		if p, ok := hash[target-x]; ok {
			return []int{p, i}
		}
		hash[x] = i
	}
	return nil
}

func ex8_test() {
	n1 := []int{2, 7, 11, 15}
	t1 := 9
	fmt.Println(twoSum(n1, t1))

	n2 := []int{3, 2, 4}
	t2 := 6
	fmt.Println(twoSum(n2, t2))

	n3 := []int{3, 3}
	t3 := 6
	fmt.Println(twoSum(n3, t3))
}

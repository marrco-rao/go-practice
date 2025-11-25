package main

import "fmt"

/*
引用类型：切片
26. 给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。
考虑 nums 的唯一元素的数量为 k。去重后，返回唯一元素的数量 k。
nums 的前 k 个元素应包含 排序后 的唯一数字。下标 k - 1 之后的剩余元素可以忽略。

不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
提示：可以使用双指针法，一个慢指针 i 用于记录不重复元素的位置，一个快指针 j 用于遍历数组，
*/
func removeDuplicates(nums []int) ([]int, int) {
	if len(nums) < 2 {
		return nums, len(nums)
	}
	slow := 0
	// 	fast 依次往后走，知道走完全部元素
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			// 找到不同元素，调整slow，移动当前fast元素进行填充
			slow++
			nums[slow] = nums[fast]
		}
	}

	return nums[:slow+1], slow + 1
}

func ex6_test() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//nums := []int{1, 1, 1, 2, 2, 2, 3, 3, 3}
	//nums := []int{1, 1, 1}
	v1, l1 := removeDuplicates(nums) // 调用
	fmt.Println("v1:", v1, "l1:", l1)
}

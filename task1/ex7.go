package main

import (
	"fmt"
	"sort"
)

/*
引用类型：切片
56. 合并区间：以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，将当前区间与切片中最后一个区间进行比较，
如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。

示例 1：
输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

示例 2：
输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。

示例 3：
输入：intervals = [[4,7],[1,4]]
输出：[[1,7]]
解释：区间 [1,4] 和 [4,7] 可被视为重叠区间。
*/

func merge(intervals [][]int) [][]int {
	// 排序: 按start值升序排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var rs [][]int
	for i := 0; i < len(intervals); i++ {
		// 合并操作
		m := len(rs)
		if m > 0 && intervals[i][0] <= rs[m-1][1] {
			if intervals[i][1] > rs[m-1][1] {
				rs[m-1][1] = intervals[i][1]
			}
		} else {
			rs = append(rs, intervals[i])
		}
	}
	return rs
}
func ex7_test() {
	i1 := [][]int{[]int{1, 3}, []int{2, 6}, []int{8, 10}, []int{15, 18}}
	i2 := [][]int{[]int{1, 4}, []int{4, 5}}
	i3 := [][]int{[]int{4, 7}, []int{1, 4}}
	i4 := [][]int{[]int{1, 4}, []int{2, 3}}

	fmt.Println(merge(i1))
	fmt.Println(merge(i2))
	fmt.Println(merge(i3))
	fmt.Println(merge(i4))
}

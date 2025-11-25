package main

import "fmt"

/*
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。
*/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		flag := true
		for j := 1; j < len(strs); j++ {
			if len(strs[j])-1 < i || strs[0][i] != strs[j][i] {
				flag = false
				break
			}
		}
		if !flag {
			return strs[0][:i]
		}
	}
	return strs[0]
}

func ex4_test() {
	strs := []string{"flower", "flow", "flight", "fly"}
	//strs1 := []string{"dog", "racecar", "car"}
	rs := longestCommonPrefix(strs)
	fmt.Println(rs)
}

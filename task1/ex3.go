package main

import (
	"fmt"
)

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。
*/

// 思路：用栈的思路来实现
func isValid(s string) bool {
	if s == "" || len(s)%2 > 0 {
		return false
	}
	// base_char和pairs 改进输入方式后能做的更灵活
	pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}
	res := []rune{}
	for _, v := range s {
		// 当前元素是不是右括号,压栈
		if pairs[v] == 0 {
			res = append(res, v)
		} else {
			// 栈里没有匹配项，直接判断失败，返回false
			if len(res) == 0 {
				return false
			}
			if res[len(res)-1] == pairs[v] {
				// 出栈
				res = res[:len(res)-1]
			} else {
				// 栈端的元素跟当前元素不匹配，判断失败，返回false
				return false
			}
		}
	}
	fmt.Println(string(res))
	if len(res) > 0 {
		return false
	}
	return true
}
func ex3_test() {
	str1 := "()"     //true
	str2 := "()[]{}" //true
	str3 := "(]"     //false
	str4 := "([])"   //true
	str5 := "([))"   //false
	str6 := "["      //false

	fmt.Println(isValid(str1))

	fmt.Println(isValid(str2))
	fmt.Println(isValid(str3))
	fmt.Println(isValid(str4))
	fmt.Println(isValid(str5))
	fmt.Println(isValid(str6))
}

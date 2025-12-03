package main

import "fmt"

/*
1.题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
考察点 ：指针的使用、值传递与引用传递的区别。
*/
func pointParamAdd(a *int) {

	if a == nil {
		fmt.Println("错误：参数为nil")
		return
	}
	fmt.Println("*a-before:", a)
	*a += 10
	//*a = *a + 10
	fmt.Println("*a-after:", a)
}
func ex1_test1() {
	num := 1
	fmt.Println("调用前:", num)
	pointParamAdd(&num)
	fmt.Println("调用后:", num)

	fmt.Println()
	var nilPtr *int // 空指针
	pointParamAdd(nilPtr)
}

/*
2.题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
考察点 ：指针运算、切片操作。
*/
func pointParamMuiltiply(b *[]int) {
	if b == nil || *b == nil {
		fmt.Println("错误：参数为nil")
	}
	//for i := 0; i < len(*b); i++ {
	//	(*b)[i] *= 2
	//}
	for i := range *b {
		(*b)[i] *= 2
	}
}
func ex1_test2() {

	var b1 []int
	fmt.Println("调用前:", b1)
	pointParamMuiltiply(&b1)
	fmt.Println("调用后:", b1)

	fmt.Println()
	b := []int{1, 0, 3}
	fmt.Println("调用前:", b)
	pointParamMuiltiply(&b)
	fmt.Println("调用后:", b)

}

//func main() {
//	ex1_test1()
//	fmt.Println("========================")
//	ex1_test2()
//}

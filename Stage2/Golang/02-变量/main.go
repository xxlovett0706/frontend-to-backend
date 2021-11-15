/*
	! 注意事项
	* 函数以外每个语句都必须以关键字开始（var、const或func）
	* :=不能用于函数外
	* _多用于占位符，表示忽略值
*/
package main

import "fmt"

// 全局变量
var globalName = "全局名字"
var globalAge = 12
var globalIsOk = true

func getNameAndAge() (string, int) {
	return "MasK", 18
}

func main() {
	// 标准声明格式
	var name string
	var age int
	var isOk bool
	fmt.Println(name, age, isOk)

	// 批量声明变量
	var (
		name1 string
		age1  int
		isOk1 bool
	)
	fmt.Println(name1, age1, isOk1)

	// 声明变量的同时指定初始值
	var name2 string = "小王子"
	var age2 int = 18
	var isOk2 bool = true
	fmt.Println(name2, age2, isOk2)

	// 类型推导
	var name3 = "小公主"
	var age3 = 12
	var isOk3 = true
	fmt.Println(name3, age3, isOk3)

	// 同时声明多个变量并赋值(有类型推导)
	var name4, age4, isOk4 = "沙赫纳扎", 16, true
	fmt.Println(name4, age4, isOk4)

	// 短变量声明(只能在函数内部声明)
	name5 := "切糕牛"
	age5 := 28
	isOk5 := false
	fmt.Println(name5, age5, isOk5)

	// 匿名变量
	name6, _ := getNameAndAge()
	_, age6 := getNameAndAge()
	fmt.Println(name6, age6)
}

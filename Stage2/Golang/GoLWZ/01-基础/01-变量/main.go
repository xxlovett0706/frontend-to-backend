// 变量
package main

import "fmt"

// 7. 声明全局变量并初始化
var school = "bat"

func main() {
	// 1. 标准声明格式
	var name string
	var age int
	var isOk bool
	fmt.Println(name, age, isOk)

	// 2. 批量声明格式
	var (
		a string
		b int
		c bool
		d float64
	)
	fmt.Println(a, b, c, d)

	// 3. 声明变量同时指定初始值
	var name1 string = "小王子"
	var age1 int = 18
	fmt.Println(name1, age1)

	// 4. 类型推导
	var name2 = "小公主"
	var age2 = 18
	fmt.Println(name2, age2)

	// 5.多变量声明并类型推导
	var name3, age3 = "沙河娜扎", 20
	fmt.Println(name3, age3)

	// 6. 短变量声明
	// ! 只能用于函数中
	name4 := "小红帽"
	age4, isChild := 8, true
	fmt.Println(name4, age4, isChild)

	// 8. 匿名变量
	res1, _ := add(3, 4)
	_, ok := add(3, -4)
	fmt.Println(res1, ok)
}

// 用于匿名变量的展示
func add(x int, y int) (int, bool) {
	result := x + y
	if result >= 0 {
		return result, true
	} else {
		return result, false
	}
}

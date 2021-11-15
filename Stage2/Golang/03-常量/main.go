/*
	! 注意事项
	* 常量定义时必须赋值
	* iota 计数器只能在常量表达式中使用；使用 iota能简化定义，在定义枚举时很有用; iota 遇到 const 则重置为 0；iota 每新增一行则加1
*/
package main

import "fmt"

// 单个声明
const PI = 3.14
const E = 2.7

// 批量声明
const (
	MAX = 10
	MIN = 0
)

// 批量申明同一个值，可省略
const (
	n1 = 10
	n2
	n3
)

// iota 常量计数器
const (
	b1 = iota // 0
	b2 = iota
	b3 = iota
)

// iota 省略某值
const (
	c1 = iota
	c2
	_
	c4
)

// iota 中间插队
const (
	a1 = iota
	a2 = iota
	a3 = 100
	a4 = iota
)
const a5 = iota

// iota 定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

// iota 多个赋值
const (
	a, b = iota + 1, iota + 2
	c, d
	e, f
)

func main() {
	fmt.Println(n1, n2, n3)
	fmt.Println(b1, b2, b3)
	fmt.Println(c1, c2, c4)
	fmt.Println(a1, a2, a3, a4, a5)
	fmt.Println(KB, MB, GB, TB, PB)
	fmt.Println(a, b, c, d, e, f)
}

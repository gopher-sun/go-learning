// Package main 演示基于自定义类型实现错误码。
package main

import "fmt"

// code 是基于 int 的自定义类型，用作错误码。
//
// 自定义类型（type X Y）与类型别名（type X = Y）的区别：
//   - 自定义类型：X 是全新的类型，拥有独立的方法集
//   - 类型别名：X 只是 Y 的另一个名字，完全等价
type code int

const (
	Scode  code = 0
	Secode code = 1001
	Necode code = 1002
)

// Getcode 为 code 类型定义方法，返回错误码对应的描述。
//
// 接收者用值类型（c code）还是指针类型（c *code）？
//   - 方法不修改接收者，且类型很小（int）→ 用值类型
//   - 需修改接收者，或类型较大 → 用指针类型
func (c code) Getcode() string {
	switch c {
	case Scode:
		return "成功"
	case Secode:
		return "服务错误"
	case Necode:
		return "网络错误"
	default:
		return "未知错误"
	}
}

func main() {
	// 逐个演示各错误码
	codes := []code{Scode, Secode, Necode, code(9999)}
	for _, c := range codes {
		fmt.Printf("错误码 %d -> %s\n", c, c.Getcode())
	}

	// 演示自定义类型的类型安全特性：
	// 下面这行如果取消注释会编译报错，因为 int 常量不能直接赋给 code
	// var x code = 1001   // 错误：cannot use 1001 (untyped int constant) as code
	var x code = code(1001) // 正确：需要显式转换
	fmt.Println("显式转换后:", x.Getcode())
}

// Package main 演示 Go 的 init 函数。
package main

import "fmt"

// 包级变量在 init 之前初始化。
var appName = "go-learning"

// init 函数的特点：
//  1. 每个文件可以有多个 init 函数，按声明顺序执行
//  2. 自动执行，不能被显式调用
//  3. 执行时机：包级变量初始化之后 → main 之前
//  4. 常用于：配置加载、注册驱动、建立连接等一次性准备工作
func init() {
	fmt.Printf("[init-1] 初始化中，应用名: %s\n", appName)
}

func init() {
	fmt.Println("[init-2] 第二个 init 函数，同样会自动执行")
}

func main() {
	fmt.Println("[main] 主函数开始执行")
}

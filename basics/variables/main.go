package main

import (
	"fmt"
	"github.com/gopher-sun/go-learning/pkg/version"
)

func hello() {
	fmt.Print("hello")
}

// 全局变量,不能用简短声明
var age = 11

// 另外一种声明方法
var (
	b1 int = 18
	b2 int = 19
)

// 声明常量
const version2 = 111

func main() {
	//声明然后赋值
	var name string
	name = "乐乐"
	fmt.Println(name)
	//声明且赋值
	var name1 string = "乐乐"
	fmt.Println(name1)
	//省略类型
	var name2 = "乐乐"
	fmt.Println(name2)
	//声明并赋值，短声明
	name3 := "乐乐"
	fmt.Println(name3)

	hello()
	fmt.Println(age)
	var a1, a2 = 13, 24
	fmt.Println(a1, a2)
	fmt.Println(version.Version1)
}

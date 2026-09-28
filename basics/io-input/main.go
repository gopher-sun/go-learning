package main

import "fmt"

func main() {
	fmt.Println("请输入您的名字：")
	var name string
	fmt.Scan(&name)
	fmt.Println(name)
	fmt.Println("请输入您的年龄：")
	var age int
	//接入一个异常类，防止类型输入错误
	n, err := fmt.Scan(&age)
	fmt.Println(n, err, age)
}

package main

import "fmt"

func main() {
	//和其他语言一样我就不写了
	var age int
	fmt.Scan(&age)
	if age > 18 {
		fmt.Println("成年")
	}
}

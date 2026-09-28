package main

import "fmt"

func main() {
	fmt.Printf("%s222\n", "乐乐")
	fmt.Printf("%d", 222)
	fmt.Printf("%T", 444)    //%T是打印类型的
	fmt.Printf("%v\n", "pp") //任意的什么都能显示出来
	fmt.Printf("%#v", "")    //可以打印空字符串
	//将内容格式化赋值，然后再打印
	var f = fmt.Sprintf("%d", 222)

	fmt.Println(f)
}

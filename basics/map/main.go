package main

import (
	"fmt"
)

func main() {
	var userMap map[int]string = map[int]string{
		1: "ll",
		2: "kk",
		3: "qq",
	}
	fmt.Println(userMap[1])
	fmt.Printf("%#v\n", userMap[3]) //确定是否有这个数据

	var name map[int]string = map[int]string{}
	name[1] = "lezi" //两种不同的赋值方法

	delete(name, 1) //通过删除key来删除
	map1, ok := name[5]
	fmt.Println(map1, ok) //通过这个布尔值确定有没有取到值
}

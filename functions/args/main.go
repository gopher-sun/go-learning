package main

import "fmt"

func copy(name string) {
	fmt.Printf("%p\n", &name) //这里会开辟一个新地址
}
func set(name *string) { //有点像指针，传地址
	fmt.Printf("%p\n", name) //因为已经定义为地址了，所以不用&
	*name = "冯"
	//通过内存地址找值

}

func main() {
	var name string = "乐乐"
	fmt.Printf("%p\n", &name) //传地址
	copy(name)
	set(&name)
	fmt.Println(name)
}

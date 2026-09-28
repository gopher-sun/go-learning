package main

import "fmt"

func main() {
	//var u8 uint8=255
	//uint8,最大值就是2^8-1,最小值是0
	//int8
	//正数时最大值是2^7-1,最小值是0
	//负数时最大值是-1,最小值是负的2^7（源码）
	var a byte = 'a'            //ascii码表里的字符
	fmt.Printf("%c %d\n", a, a) //%c输出单字符 %d输出acii码
	var z rune = '中'
	fmt.Printf("%c %d", z, z)
}

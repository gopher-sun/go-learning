package main

import (
	"fmt"
	//"time"
)

func main() {
	sum := 0
	for i := 1; i < 100; i++ {
		sum += i
	}
	fmt.Println(sum)
	//go中的死循环，go没有while
	/*for {
		fmt.Println(time.Second)
		time.Sleep(1*time.Second)
	}*/
	//while模式
	var p int = 1
	var sum1 int = 0
	for p <= 100 {
		sum1 += p
		p++
	}
	fmt.Println(sum1)
	//do-while模式
	var t int = 1
	var sum2 int = 0
	for {
		sum2 += t
		t++
		if t >= 100 {
			break
		}
	}
	fmt.Println(sum2)
	//列表循环
	var list = []string{"ll", "ee"}
	for i := 0; i < len(list); i++ {
		fmt.Println(list[i])
	}
	//for range循环
	for index, item := range list {
		fmt.Println(index, item)
	}
	//循环map,for range循环
	var usermap = map[string]string{"name": "qq"}
	for key, value := range usermap {
		fmt.Println(key, value)
	}
}

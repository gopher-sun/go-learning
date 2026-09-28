package main

import "fmt"

func main() {
	//第一种用法
	var age int
	fmt.Println("请输入您的名字")
	fmt.Scan(&age)
	switch {
	case age <= 0:
		fmt.Println("未出生")
		fallthrough //让他继续往下走，go的switch没有break，自动截断，不会往下走
	case age < 18:
		fmt.Println("未成年")
	default:
		fmt.Println("中年")
	}

	//-----------------------------------------------------------------------------------------------------
	//第二种用法
	var week int
	fmt.Scan(&week)
	switch week {
	case 1:
		fmt.Println("周一")
	case 2:
		fmt.Println("周二")
	case 3:
		fmt.Println("周三")
	default:
		fmt.Println("无")
	}
}

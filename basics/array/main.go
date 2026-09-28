package main

import "fmt"

func main() {
	var namelist [3]string = [3]string{"aa", "bb", "cc"}
	fmt.Println(namelist)
	fmt.Println(namelist[0])
	fmt.Println(namelist[2])
	fmt.Println(len(namelist))
	fmt.Println(namelist[len(namelist)-1])

}

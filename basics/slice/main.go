package main

import (
	"fmt"
	"sort"
)

func main() {
	//只能是同类型的
	var namelist []string
	//var namelist []string=[]string{}可以这样进行初始化
	//var namelist =[]string {}
	//namelist :=[]string{}
	//namelist =make([]string,0)
	//以上的初始化都可以
	namelist = append(namelist, "qq")
	namelist = append(namelist, "ww")
	namelist = append(namelist, "ee")
	fmt.Println(namelist[0])
	fmt.Println(namelist[1])
	//make函数
	agelist := make([]int, 3)
	fmt.Println(agelist)
	//切片排序
	array := [3]int{1, 2, 3}
	slices := array[:]
	fmt.Println(slices)
	fmt.Println(array[0:2])
	fmt.Println(array[0:1])
	var ints = []int{4, 6, 7, 9}
	//排序函数
	sort.Ints(ints)
	fmt.Println(ints) //默认升序
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	fmt.Println(ints) //降序
}

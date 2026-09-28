package main

import (
	"fmt"
)

type Class struct {
	name string
}
type Student struct {
	Class //相当于继承
	name  string
	age   int
}

func (s Student) study() {
	fmt.Printf("%s学习\n", s.name)
}
func (s Student) info() {
	fmt.Printf("名字%s 班级%s", s.name, s.Class.name)
} //继承了class里的变量
func (s *Student) setName(name string) { //这里的指针给接收的参数所以和外界指针不太一样
	s.name = name
	fmt.Printf("in%p\n", &s)
}
func main() {
	c1 := Class{
		name: "ww",
	}
	s1 := Student{Class: c1, name: "ii"}
	s1.study()
	s2 := Student{Class: c1, name: "oo", age: 11}
	s2.study()
	s2.setName("uu")
	fmt.Printf("out%p\n", &s2) //两次打印地址不一样
	s2.study()
	fmt.Printf("out%p\n", &s2) //两次打印地址不一样
}

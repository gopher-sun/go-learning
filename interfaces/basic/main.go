package main

import "fmt"

type singiterface interface {
	Sing()
	Getname() string
}

type chicken struct {
	name string
}

func (c chicken) Sing() {
	fmt.Println("ikun")
}
func sing(s singiterface) {
	s.Sing()
	fmt.Println(s.Getname())
}
func (c chicken) Getname() string {
	return c.name
}
func main() {
	var s singiterface = chicken{"小黑子"}
	s.Sing()
	sing(s)
}

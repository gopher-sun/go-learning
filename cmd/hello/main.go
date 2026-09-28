// Command hello 是一个演示程序，展示项目的标准结构。
package main

import (
	"fmt"
	"os"

	"github.com/gopher-sun/go-learning/internal/greet"
)

func main() {
	name := "gopher-sun"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	msg, err := greet.Hello(name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "错误: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(msg)
}

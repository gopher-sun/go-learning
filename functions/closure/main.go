package main

import (
	"fmt"
	"time"
)

func awaitAdd(awaitsecond int) func(...int) int {
	time.Sleep(time.Duration(awaitsecond) * time.Second)
	return func(numberList ...int) (sum int) {
		for _, i := range numberList {
			sum += i
		}
		return sum
	}
}

func main() {
	t1 := time.Now()
	fmt.Println(awaitAdd(2)(1, 2, 3))
	subTime := time.Since(t1)
	fmt.Println(subTime)
}

package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string `json:"name"`
	Age      int    `json:"age,omitempty"` //抛弃空值，当用结构体定义时如果没有定义这个字段就默认为0，然后用这个函数就不显示了
	Password string `json:"-"`             //转json时忽略该字段
}

func main() {
	user := User{
		Name: "ll", Password: "234532",
	}
	byteData, _ := json.Marshal(user)
	fmt.Println(string(byteData))
}

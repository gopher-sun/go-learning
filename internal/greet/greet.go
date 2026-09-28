// Package greet 提供问候相关的功能。
//
// 这是 internal 包的示例：internal 下的包只能被本模块内部引用，
// 外部模块无法导入，适合存放不想对外暴露的实现。
package greet

import (
	"errors"
	"fmt"
)

// ErrEmptyName 表示传入的姓名为空。
var ErrEmptyName = errors.New("name cannot be empty")

// Hello 返回对指定姓名的问候语。
//
// 如果 name 为空字符串，返回 ErrEmptyName。
func Hello(name string) (string, error) {
	if name == "" {
		return "", ErrEmptyName
	}
	return fmt.Sprintf("Hello, %s! 欢迎来到 Go 世界。", name), nil
}

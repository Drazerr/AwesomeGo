package main

import "fmt"

func main() {
	var description interface{}
	switch description.(type) {
	case string:
		fmt.Println("是个字符串")
	case int:
		fmt.Println("是个数字")
	default:
		fmt.Println("不知道诶")
	}
}

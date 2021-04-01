package main

import "fmt"

func main() {
	description := "ugly"
	switch description {
	case "beautiful":
		fmt.Println("爱死你了！")
	case "ugly":
		fmt.Println("找死吗？")
		fallthrough
	default:
		fmt.Println("我没听懂啊……")
	}
}

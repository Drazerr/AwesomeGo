package main

import "fmt"

func main() {
	description := "ugly"
	switch description {
	case "ugly":
		fmt.Println("找死吗？")
	case "beautiful":
		fmt.Println("爱死你了！")
	default:
		fmt.Println("我没听懂啊……")
	}
}

package main

import "fmt"

func main() {
	description := "ugly"
	switch description {
	case "beautiful", "cute", "handsome":
		fmt.Println("爱死你了！")
	case "ugly":
		fmt.Println("找死吗？")
	default:
		fmt.Println("我没听懂啊……")
	}
}

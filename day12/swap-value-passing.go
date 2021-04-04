package main

import "fmt"

func main() {
	/* 定义局部变量 */
	var a = "cute"
	var b = "ugly"

	fmt.Printf("Penguin is %s?\n", a)
	fmt.Printf("or Penguin is %s?\n", b)

	swap(a, b)

	fmt.Printf("Not sure if Penguin is %s?\n", a)
	fmt.Printf("Not sure if Penguin is %s either.\n", b)

	swapRef(&a, &b)

	fmt.Printf("Penguin is %s?\n", a)
	fmt.Printf("No Penguin is %s!\n", b)
}

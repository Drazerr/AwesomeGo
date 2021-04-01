package main

import "fmt"

func sortDes(a, b int) (int, int) {
	if a < b {
		return b, a
	}
	return a, b
}

func main() {
	fmt.Println(sortDes(1, 2))
}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 1; i < 20; i++ {
		if i%2 != 0 { // 跳过循环
			continue
		}
		fmt.Println("霜羽在第 " + strconv.Itoa(i) + " 次循环表示 ta 很酷")
	}
}

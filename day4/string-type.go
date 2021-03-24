package main

import "fmt"

func main() {
	question := "Hoarfroster is cute"
	response := `I agree with you`
	fmt.Println(question, response)

	stringWithQuote := "Hoarfroster love \""
	byteWithQuote := '\''
	fmt.Println(stringWithQuote, byteWithQuote)

	// 缩进
	fmt.Println("Hello \t Hoarfroster")
	// Hello	Hoarfroster

	// 输出换行
	fmt.Println("Hello \n Hoarfroster")
	// Hello
	// Hoarfroster

	// 移动光标到行首（回车）
	fmt.Println("Hello \r Hoarfroster")
	//  Hoarfroster
}

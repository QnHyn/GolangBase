package main
import "fmt"

// golang 中标识符的使用
func main() {
	//Golang中严格区分大小写
	var num int = 10
	var Num int = 20
	fmt.Println(num, Num)

	// 标识符中不能包含空格
	// _ 是空标识符，只能用于占位，不能用于变量
	var _ int = 12 //err
}
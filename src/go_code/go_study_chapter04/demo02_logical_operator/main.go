package main
import "fmt"

// 声明一个测试函数
func test() bool {
	fmt.Println("test...")
	return true
}


func main() {
	// golang 中关系运算符的 逻辑与和逻辑或
	var  i int = 10 
	if i > 20 && test(){
		// 这里根本不会执行test中的函数 短路与
		// 第一个条件为false, 结果为false.不执行test.和if中语句
		fmt.Println("ok与")
	}
	if  i > 1 || test() {
		// 这里根本不会执行test中的函数 短路或
		// 第一个条件为true 结果为true 不执行test 执行if语句
		fmt.Println("ok或")
	}
}
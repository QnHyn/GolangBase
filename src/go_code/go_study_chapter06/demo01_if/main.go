package main
import "fmt"


func main() {
	// 使用陷阱
	var age int
	fmt.Println("输入您的年龄")
	fmt.Scanln(&age)
	if age > 18 {
		fmt.Println("ok1")
	} else if age > 9 {
		fmt.Println("ok2")
	} else if age > 4 {
		fmt.Println("ok3")
	}
}
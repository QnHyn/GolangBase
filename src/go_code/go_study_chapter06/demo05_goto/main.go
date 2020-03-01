package main
import (
	"fmt"
) 


func main() {
	// 演示goto的使用
	fmt.Println("ok1") 
	goto label
	fmt.Println("ok2")
	fmt.Println("ok3")
	fmt.Println("ok4")
	label:
	fmt.Println("0k5")
	fmt.Println("ok6") 
	fmt.Println("ok7")
}
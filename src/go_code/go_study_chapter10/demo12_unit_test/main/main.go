package main

import (
	_ "fmt"
)

// 一个被测试的函数
func addUpper(n int) int{
	res := 0
	for i:=0; i <= n; i++{
		res += i
	}
	return res
}
// func main(){
// 	//传统的测试方法，就是在main函数中使用看看结果是否正确
// 	res := addUpper(10) // 1.+ 10 = 55
// 	if res != 55{
// 		fmt.Printf("addupper错误返回值=%v期望值=%v\n", res, 55)
// 	}else{
// 		fmt.Printf("addupper正确返回值=%v期望值=%v\n", res, 55)
// 	}
// }
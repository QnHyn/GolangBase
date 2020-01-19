package main
import "fmt"



func main() {

	var price float32 = 89.12
	fmt. Println("price=", price)

	//尾数部分可能丢失，造成精度损失。- 123.0000901
	var num1 float32  = -123.0000901
	var num2 float64  = -123.0000901
	fmt. Println("num1=", num1, "num2=", num2)

	// 浮点注意事项：
	// 十进制表示形式
	num3 := 5.12
	num4 := .234 //=>0.234
	fmt. Println("num3=", num3, "num4=", num4)

	//科学计数法表示e或E
	num8 := 5.1234e2 //=> 5.1234 * 10^2
	num9 := 5.1234E-2 //=> 5.1234 * 10^-2 
	fmt. Println("num8=", num8, "num9=", num9)

}
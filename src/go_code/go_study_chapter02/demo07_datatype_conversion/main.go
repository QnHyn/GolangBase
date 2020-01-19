package main
import "fmt"


// 演示golang中基本数据类型的转换
func main() {
	var i int = 99
	// 把i转换为float类型。99.000000类型是:float32,99的数据类型:int
	var n1 float32 = float32(i)
	fmt.Printf("%f类型是:%T, \n%v的数据类型:%T", n1, n1, i, i)

	// 溢出的处理 int64 -> int8。具体溢出的结果不确定，反正想不到的结果
	var num int64 = 999998
	var num1 int8 = int8(num)
	fmt.Println("num1=", num1)
}
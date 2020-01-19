package main
import(
	"fmt"
) 


// 演示golang中复杂数据类型 指针
func main() {
	// 基本数据类型在内存的布局
	var i int = 10
	// i变量的内存地址是多少呢？ &i
	fmt.Println("i的内存地址:",&i)

	// 下面的 var ptr *int = &i
	// 1. ptr是一个指针变量。类型为*int。ptr本身的值是&i
	var ptr *int = &i
	fmt.Printf("ptr=%v\n", ptr)

	// 2. 取出指针存的地址指向的值
	fmt.Printf("*ptr指向的值:%v\n", *ptr)

	// 3. 通过ptr给i赋值 i的值:20
	*ptr = 20
	fmt.Printf("i的值:%v\n", i)

}
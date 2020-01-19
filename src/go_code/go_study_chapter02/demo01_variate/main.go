package main
import "fmt"

func main() {
	//定义变量/声明变量
	var i int
	// 给 i 赋值
	i = 10
	// 使用变量
	fmt.Println("i=", i)
	// +的两种使用
	var k, p int
	var s1, s2 = "hello", "world"
	p = k + i
	var s = s1 + s2
	fmt.Println("p=", p)
	fmt.Println("s=", s)
}
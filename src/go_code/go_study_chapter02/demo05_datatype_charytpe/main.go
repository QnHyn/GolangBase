package main
import "fmt"


// 演示golang中字符类型的使用
func main() {
	var c1 byte = 'a'
	var c2 byte = '0'

	// 我们直接输出，实际上输出的是字符对应的ASCII码值
	fmt.Println("c1:", c1, "\nc2:", c2)
	// 如果想要输出对应的的字符的话，需要使用格式化的输出方式
	fmt.Printf("c1=%c c2=%c\n", c1, c2)
	
	//var c3 byte = '北'  constant 21271 overflows byte
	var c3 int = '北'
	fmt.Printf("c3=%c\nc3对应的码值%d", c3, c3)

}
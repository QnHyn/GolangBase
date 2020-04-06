package main
import "fmt"

//累加器
func AddUpper() func (int) int {
	var n int=10
	var str string = "hello"
	return func (x int) int {
		n = n + x
		str += str + string(x)
		fmt.Println("str=", str)
		return n 
	}
}

func main() {

	//使用前面的代码
	f := AddUpper()
	fmt.Println(f(1))// 11
	// 这里需要特别注意 上面的值11 又赋值给了n 11+2=13
	fmt.Println(f(2))// 13
	fmt.Println(f(3))// 16
}
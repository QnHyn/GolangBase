package main
import "fmt"


func main() {
	// golang 中位运算和位移运算
	fmt.Println(2&3) //2
	fmt.Println(2|3) //3
	fmt.Println(2^3) //1

	fmt.Println(-2&2) //2
	fmt.Println(-2|2)//-2
	fmt.Println(-2^2)//-4

	a := 1 >> 2 //0
	b := -1 >> 2 //-1
	c := 1 << 2 // 4
	d := -1 << 2 //-4
	fmt.Println(a, b, c, d)
 }
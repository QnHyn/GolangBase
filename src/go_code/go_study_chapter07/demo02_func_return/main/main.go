package main
import (
	"fmt"
)


func getSumAndsub(n1 int, n2 int) (int, int) {
	sum:=n1+n2
	sub:=n1-n2
	return sum, sub
}

func main() {
	// 忽略返回的减法的结果
	sum, _ := getSumAndsub(10, 20)
	fmt.Println("main sum =", sum) //30
	//调用getSumAndsub
	res1, res2 := getSumAndsub(1, 2) //res1 = 3 res2 = -1 
	fmt.Printf("res1=%v res2=%v\n", res1, res2)
}
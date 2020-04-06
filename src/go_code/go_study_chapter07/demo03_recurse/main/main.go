package main
import (
	"fmt"
)


func fib(n int) int {
	if(n==1 || n==2){
		return 1
	}else{
		return fib(n-1) + fib(n-2)
	}
}

func main() {
	res := fib(3)
	//测试
	fmt.Println("res=", res)
	fmt.Println("res=", fib(4))
	fmt.Println("res=", fib(5))
	fmt.Println("res=", fib(6))
}
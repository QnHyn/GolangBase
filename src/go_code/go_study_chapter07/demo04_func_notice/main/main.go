package main
import (
	"fmt"
)


//在Go中， 函数也是一种数据类型，
//可以赋值给一个变量，则该变量就是一个函数类型的变量了。通过该变量可以对函数调用
func getSum(n1 int, n2 int) int {
	return n1 + n2
}

//函数既然是一种数据类型，因此在Go中，函数可以作为形参，并且调用
func myFun(funvar func(int, int) int, num1 int, num2 int ) int {
	return funvar (num1, num2)
}

//支持对函数返回值命名, 返回不需要规定位置啦
func getSumAndsub(n1 int, n2 int) (sum int, sub int){
	sub=n1-n2
	sum=n1+n2
	return
}

//案例演示:编写一个函数sum , 可以求出1到多个int的和
//可以参数的使用
func sum(n1 int,args... int) int {
	sum := n1
	//遍历args
	for i := 0; i < len(args); i++ {
	sum += args[i] //args[0] 表示取出args切片的第一个元素值，其它依次类推
	}
	return sum
}

func main() {
	a := getSum
	fmt.Printf("a的类型%T, getsum类型是%T\n", a, getSum)
	res := a(10, 40) //等价res := getsum(10, 40) 
	fmt.Println("res=", res)
	
	res2 := myFun(getSum, 50, 60)
	fmt.Println("res2=", res2)
	
	//给int取了别名，在go中myInt 和int 虽然都是int类型， 但是go认为myInt和int两个类型
	type myInt int
	var num1 myInt 
	var num2 int
	num1 = 40
	num2 = int(num1) //各位， 注意这里依然需要显示转换，go认为myInt和int两个类型
	fmt.Println("num1=" , num1, " num2=" , num2 )

	a1, b1 := getSumAndsub(1, 2)
	fmt.Printf("a=%v b=%v\n", a1, b1)

	//测试一下可变参数的使用
	res4 := sum(10, 0, -1, 90, 10,100)
	fmt.Println("res4=", res4)
}
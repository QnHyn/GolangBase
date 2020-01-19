package main
import "fmt"

// 定义全局变量
// var  p1 = 100
// var  p2 = 200
// var  pname = "jack"
//var p1, p2, pname = 100, 200, "jack"
//(这种赋值不能定义全局变量)p1, p2, pname := 100, 200, "jack"
// 全局变量的name, 会被mian函数中的name赋值替代
var(
	p1 = 100
	p2 = 200
	name = "jack"
)




func main() {
	//变量的使用方式一
	//第一种:指定变量类型，声明后若不赋值，使用默认值
	// int 的默认值为0，其他变量的默认值看数据类型章节
	var i int
	fmt.Println("i=", i)

	//变量的使用方式二
	//第二种:根据值自行判定变量类型(类型推导)。动态编程语言的特点python,php
	// go语言是快速的、静态类型的编译型语言，感觉却像动态类型的解释型语言。
	var num = 10.11
	fmt.Println("num=", num)

	//变量的使用方式三
	//第三种:省略var, 注意:=左侧的变量不应该是已经声明过的，否则会导致编译错误(:绝对不能省掉)
	// 等价于 var name string  name = "tom"
	name := "tom"
	fmt.Println("name=", name)

	//多变量声明方式一(一次声明多个变量)
	var n1, n2, n3 int
	fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)

	//多变量声明方式二(一次声明多个不同类型变量)
	var n4, name1, n5 = 100, "tom", 888
	fmt.Println("n4=", n4, "name1=", name1, "n5=", n5)

	//多变量声明方式三(一次声明多个不同类型变量)
	n7, name2, n8 := 100, "xiaoming", 888
	fmt.Println("n7=", n7, "name2=", name2, "n8=", n8)

	// 在go函数外部定义的变量就是全局变量
	fmt.Println("p1=", p1, "p2=", p2, "name=", name)

}
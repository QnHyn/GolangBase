package main
import "fmt"

// golang 中运算符的使用 重点 除法 取模 自增自减
func main() {
	// 说明参与运算的都是整数，除法后会截取掉小数点部分保留整数部分
	fmt.Println(10 / 4) //2
	var n1 float32 = 10 / 4 //2
	fmt.Println(n1)
	// 如果希望保留小数部分，需要有浮点数参与运算
	var n2 float32 = 10.0 / 4 //2.5
	fmt.Println(n2)
	
	// 取模 % 符号和被除数的符号一致
	fmt.Println(10 % 3)
	fmt.Println(-10 % 3)
	fmt.Println(10 % -3)
	fmt.Println(-10 % -3)

	// 自增和自减
	var i int = 10;
	i++ // i + 1
	fmt.Println(i) 
	i-- // i - 1
	fmt.Println(i)
	
	//在golang中++和--只能独立使用，不能赋值,不能赋值后比较
	//var a int;
	//a = i++ 会报错
	//if i++ > 0 {} 也会报错
	// ++i ++和-- 不能写在前面


	// 练习一:假如还有97天放假，问: xx个星期零xx天
	var days int = 97
	var week int = days / 7
	var day int = week % 7
	fmt.Printf("还剩下%d个星期零%d天\n", week, day)

	// 练习二:定义一个变量保存华氏温度，华氏温度转换摄氏温度的公式为:5/9* (华氏温度- 100),请求出华氏温度对应的摄氏温度 I
	var huashi float32 = 134.2
	var wenshi = 5.0 / 9 * (huashi - 100)
	fmt.Printf("华氏温度%v对应的摄氏温度%v\n", huashi, wenshi)
}
package main
import(
	"fmt"
	"strconv"
) 


// 演示golang中string类型转为其他基本数据类型
func main() {
	// string类型转为布尔类型
	var str string = "true"
	var b bool
	// 1. strconv.ParseBool(str)函数会返回两个值(value bool, err error)
	// 2.因为我只想获取到value bool,不想获取err所以我使用_忽略
	b, _ = strconv.ParseBool(str)
	fmt.Printf("b type %T b=%v\n", b, b)
	
	// string类型转为整数(默认返回的是int64)，如果需要变成int32.转一下
	var str2 string = "1234590"
	var n1 int64
	n1, _ = strconv.ParseInt(str2, 10, 64) //十进制，64位
	fmt.Printf("b type %T b=%v\n", n1, n1)

	// string类型转为float类型(默认返回的是float64).如果需要变成float32.转一下
	var str3 string = "123.456"
	var f1 float64
	f1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("f1 type %T f1=%v\n",f1, f1)

	//注意:n3 type int64 n3=0
	var str4 string = "hello" 
	var n3 int64
	n3, _ = strconv.ParseInt(str4, 10, 64)
	fmt.Printf("n3 type %T n3=%v\n" ,n3, n3)
}
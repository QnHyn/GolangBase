package main
import(
	"fmt"
	"strconv"
) 


// 演示golang中基本数据类型转为string类型
func main() {
	var num1 int = 99
	var num2 float64 = 23.456 
	var b bool = true
	var myChar byte = 'a'
	var str1, str2, str3, str4 string //空的str
	
	//使用第一种方式来转换fmt . Sprintf方法
	str1 = fmt.Sprintf("%d", num1)
	str2 = fmt.Sprintf("%f", num2)
	str3 = fmt.Sprintf("%t", b)
	str4 = fmt.Sprintf("%c", myChar)
	fmt.Printf("str type %T str=%v\n", str1, str1)
	fmt.Printf("str type %T str=%v\n", str2, str2)
	fmt.Printf("str type %T str=%v\n", str3, str3)
	fmt.Printf("str type %T str=%v\n", str4, str4)

	//使用第二种方式来转换fmt。strconv包的函数
	str1 = strconv.FormatInt(int64(num1), 10) // 10表示十进制，2表示二进制
	str2 = strconv.FormatFloat(num2, 'f', 10, 64) //'f'表示显示格式(10进制或科学技术e),10精度,64表示num2是float64
	str3 = strconv.FormatBool(b)
	fmt.Printf("str type %T str=%v\n", str1, str1)
	fmt.Printf("str type %T str=%v\n", str2, str2)
	fmt.Printf("str type %T str=%v\n", str3, str3)

	// strconv包中的Itoa非常好用.把int转为string
	// %q,加了""  str type string str="1542"
	var num5 int = 4567
	var str5 string = strconv.Itoa(num5)
	var num6 int64 = 1542
	var str6 string = strconv.Itoa(int(num6))
	fmt.Printf("str type %T str=%v\n", str5, str5)
	fmt.Printf("str type %T str=%q\n", str6, str6)
}
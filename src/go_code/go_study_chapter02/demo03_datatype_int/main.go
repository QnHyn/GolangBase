package main
// import "fmt"
// import "unsafe"
import(
	"fmt"
	"unsafe"
)


func main() {
	//有符号整数类型 int8 -128~127
	// int16, int32, int64类推。-2^(n-1) ~ (2^(n-1)) - 1
	var i int8 = 1
	fmt.Println("i=", i)

	// var a int8 = -129 报错 constant -129 overflows int8
	var a int8 = -128
	fmt.Println("a=", a)

	//无符号整数类型。前面加u。如uint8, uint16, uint32, uint64
	//数值范围 0~(2^n)-1
	var k uint8 = 255
	fmt.Println("k=", k)

	// int, uint跟操作系统的位数有关
	// rune, byte(存放字母) 
	var b int = 8900
	fmt.Println("b=", b)

	var b1 byte = 255
	fmt.Println("b1=", b1)

	//整型的使用细节
	var n = 100
	// 如何查看变量的数据类型呢？fmt.Printf() 用于格式化输出
	fmt.Printf("n 的数据类型 %T \n", n)
	// 如何在程序查看某个变量的字节大小和数据类型
	var n1 int64 = 100
	// unsafe包的函数,可以返回n1变量占用的字节数
	fmt.Printf("n1 的数据类型 %T \n n1占用的字节大小 %d", n1, unsafe.Sizeof(n1))

}
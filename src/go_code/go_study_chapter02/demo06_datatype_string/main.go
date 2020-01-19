package main
import "fmt"


// 演示golang中字符串类型的使用
func main() {
	// string 的基本使用
	var address string = "北京上海 杭州武汉"
	fmt.Println(address)

	// 字符串一旦赋值了，字符串就不能修改了:在**Go中字符串是不可变的**。
	//str[0] = 'a' 不能修改string的内容 cannot assign to str[0]go
	var str string = "hello"
	fmt.Println(str)
	//双引号,会识别转义字符
	//反引号，以字符串的原生形式输出，包括换行和特殊字符，可以实现防止攻击、输出源代码等效果[ 案例演示]
	str2 := "abc\nabc"
	fmt.Println(str2)
	str3 := `
		package main
		import "fmt"
		
		func main() {
			//定义变量/声明变量
			var i int
			// 给 i 赋值
			i = 10
			// 使用变量
			fmt.Println("i=", i)
			// +的两种使用
			var k, p int
			var s1, s2 = "hello", "world"
			p = k + i
			var s = s1 + s2
			fmt.Println("p=", p)
			fmt.Println("s=", s)
		}
	`
	fmt.Println(str3)

	// 字符串的拼接
	var str4 string = "world"  + "golang"
	str4 += "world"
	fmt.Println(str4)

	//如果拼接的字符串很长，可以分行写。加号+一定要保留在上面行
	var str5 = "world"  + "golang" + "world"  + "golang" +"world"  + "golang" +
	"world"  + "golang" + "world"  + "golang" + "world"  + "golang" +"world"  + 
	"world"  + "golang"
	fmt.Println(str5)

	var a int//0
	var b float32//0
    var c float64//0
    var isMarried bool // false
	var name string //""
	// v表示按照变量的值输出
    fmt.Printf(" a=%d , b=%f , c=%f，isMarried=%v name=%v" ,a,b,c, isMarried, name)
}
package utils
import "fmt"

//封装成函数,方便调用
//为了让其它包的文件，可以访问到本包的函数，则该函数名的首字母需要大写，类似其它语言的public ,这样才能跨包访问。比如utils.go
func Cal(n1 float64, n2 float64, operater byte) float64{
	//输入两个数，再输入一个运算符(+,-.*,/),得到结果.。
	var res float64 

	switch operater {
		case '+':
			res = n1 + n2
		case '-':
			res = n1 - n2
		case '*':
			res = n1 * n2
		case '/':
			res = n1/n2
		default:
			fmt.Println("运算符不合法！！！")
	}
	return res
}
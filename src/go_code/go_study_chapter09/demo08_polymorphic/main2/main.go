package main
import (
	"fmt"
)

type student struct{
	Name string
}
//编写一个函数，可以判断输入的参数是什么类型
// ...可变参数(可接收任意多个任意类型的参数)
func TypeJudge(items... interface{}) {

	for index, x := range items {
		switch x.(type) {
			case bool:
				fmt.Printf("第%v个参数是bool 类型, 值是%v\n", index, x)
			case float32:
				fmt.Printf("第%v个参数是float32 类型, 值是%v\n", index, x)
			case float64:
				fmt.Printf("第%v个参数是float64 类型, 值是%y\n", index, x)
			case int, int32, int64 :
				fmt.Printf("第%v个参数是整数类型, 值是%v\n", index, x)
			case string :
				fmt.Printf("第%v个参数是string 类型, 值是%v\n", index, x)
			case student :
				fmt.Printf("第%v个参数是Student类型, 值是%v\n", index, x)
			case *student :
				fmt.Printf("第%v个参数是*student 类型, 值是%v\n", index, x)
			default :
				fmt.Printf("第%v个参数是类型 不确定, 值是%v\n", index, x)
		}
	}
}

func main() {
	var n1 float32 = 1.1
	var n2 float64 = 2.3
	var n3 int32 = 30
	var name string = "tom"
	address := "北京"
	n4 := 300
	stu := student{"xiaoming"}
	TypeJudge(n1, n2, n3, name, address, n4, stu, &stu)
}
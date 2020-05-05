package main
import (
	"fmt"
)

type Stu struct{
	Name string
	Age int
}

func main() {
	//在创建结构体变量时，就直接指定字段的值
	var stu1 = Stu{"小明", 19}
	stu2 := Stu{"小明~把", 29}

	//在创建结构体变量时，把字段名和字段值写在一起，这种写法，就不依赖字段的定义顺序.
	var stu3 = Stu{
		Name : "jack",
		Age : 13,
	}
	stu4 := Stu{
		Name : "Tony",
		Age : 15,
	}
	fmt.Println(stu1, stu2, stu3, stu4)

	//方式2, 返回结构体的指针类型在main函数中。
	//在创建结构体指针变量时，把字段名和字段值写在一起，这种写法，就不依赖字段的定义顺序
	var stu5 *Stu = &Stu{"小王", 29}
	stu6 := &Stu{"小紫", 21}
	stu7 := &Stu{
		Age: 23,
		Name: "小杨",
	}
	var stu8 = &Stu{
		Name: "小罗",
		Age: 39,
	}
	fmt.Println(stu5, stu6, stu7, stu8)
}

package main
import (
	"fmt"
)


type Monster struct{
	Name string
	Age int
}

type E struct{
	Monster
	int // 匿名字段时基本数据类型
	n int
}

func main(){
	//演示一下匿名字段时基本数据类型的使用
	var e E
	e.Name = "狐狸精"
	e.Age = 300
	e.int = 20
	e.n = 40
	fmt.Println("e=", e)
}

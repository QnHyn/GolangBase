package main
import (
	"fmt"
)

type A struct{
	Name string
	Age int
}

func (a *A) SayOk(){
	fmt.Println("A SayOk", a.Name)
}

func (a *A) hello(){
	fmt.Println("A hello", a.Name)
}

type B struct{
	A
	Name string
}

type C struct{
	Name string
	Score float64
}

type D struct{
	A 
	C
}

type F struct{
	a A // 有名结构体 组合关系
}

func main() {
	// 1. 通过B结构体访问 A结构体的方法 小写或者大写
	var b B
	b.A.Name = "tom"
	b.A.Age = 12
	b.A.SayOk()
	b.A.hello()
	// 2. 简化访问方式
	b.SayOk()
	b.hello()
	// 3. 就近访问原则 如果B有 直接访问 当然也可以指定访问
	b.Name = "xiao"
	fmt.Println(b.Name)
	fmt.Println(b.A.Name)
	// 4. 父类中有相同字段 且子类没有字段 那么必须指定访问那个父类的字段 否则会报错
	var d D
	// d.Name = "ming"
	// fmt.Println(d.Name)
	d.A.Name = "ming"
	fmt.Println(d.A.Name)

	//如果D中是一个有名结构体，则访问有名结构体的字段时，就必须带上有名结构体的名字
	var f F
	f.a.Name = "小军"
	fmt.Println(f.a.Name)
}

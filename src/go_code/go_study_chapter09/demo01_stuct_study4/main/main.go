package main
import (
	"fmt"
	"encoding/json"
)

type A struct {
	Num int
}

type B struct {
	Num int
}

// `json:"name"` 中间不能有空格呀 `json: "name"`
type Monster struct{
	Name string `json:"name"`  //`json:"name"`就是struct tag
	Age int `json:"age"`
	Skill string `json:"skill"`
}

func main() {
	var a A
	var b B
	a = A(b) // ?可以转换，但是有要求，就是结构体的的字段要完全一样(包括:名字、 个数和类型!
	fmt.Println(a, b)

	//1.创建一个Monster变量
	monster := Monster{"牛魔王", 500, "芭蕉扇~"}
	fmt.Println(monster)
	//2.将monster变量序列化为json格式字串
	// json.Marshal函数中使用反射，这个讲解反射时，我会详细介绍
	jsonstr, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json处理错误", err) 
	}
	fmt.Println("jsonstr", string(jsonstr))
}

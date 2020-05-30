package main

import (
	"fmt"
	"encoding/json"
)

type Monster struct{
	Name string
	Age int
	Birthday string
	Sal float64
	Skill string
}

// 结构体序列化
func testStruct(){
	monster := Monster{
		Name :"牛魔王",
		Age : 500,
		Birthday : "2011-11-11",
		Sal : 8000.0,
		Skill : "牛魔拳",
	}
	// 将monster序列化
	data, err := json.Marshal(&monster)
	if err != nil{
		fmt.Printf("序列化失败:%v\n", err)
	}
	fmt.Printf("机构体monster序列化后结果:%v\n", string(data))
}

//将map进行序列化
func testMap(){
	//定义一个map
	var a map[string]interface{}
	//使用map,需要make
	a = make(map[string]interface{})
	a["name"]= "红孩儿"
	a["age"]= 30
	a["address"] = "洪崖洞"
	// 将Map序列化
	data, err := json.Marshal(a)
	if err != nil{
		fmt.Printf("序列化失败:%v\n", err)
	}
	fmt.Printf("Map a序列化后结果:%v\n", string(data))
}

//演示对切片进行序列化,我们这个切片[]map[string]interface{}
func testSlice(){
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	//使用map前，需要先make
	m1["name"] = "jack"
	m1["age"] = "7"
	m1["address"] = "北京"
	slice = append(slice, m1)

	var m2 map[string]interface{}
	//使用map前, 需要先make
	m2 = make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"]= "20"
	m2["address"] = [2]string{"墨西哥","夏威夷"}
	slice = append(slice, m2)

	// 将切片序列化
	data, err := json.Marshal(slice)
	if err != nil{
		fmt.Printf("序列化失败:%v\n", err)
	}
	fmt.Printf("Slice序列化后结果:%v\n", string(data))
}

//对基本数据类型序列化，对基本数据类型进行序列化意义不大
func testFloat64(){
	var numl float64 = 2345.67
	//对numl进行序列化
	data, err := json.Marshal(numl)
	if err != nil{
		fmt.Printf("序列化失败:%v\n", err)
	}
	fmt.Printf("float64序列化后结果:%v\n", string(data))
}
func main(){
	testStruct()
	testMap()
	testSlice()
	testFloat64()
}
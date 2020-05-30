package main

import (
	"fmt"
	"encoding/json"
)

type Monster struct{
	Name string `json:"monster_name"`
	Age int `json:"monster_age"`
	Birthday string
	Sal float64
	Skill string
}

// 结构体反序列化
func unmarshalStruct(){
	// 说明str在项目开发中，是通过网络传输获取到..或者是读取文件获取到
	str := "{\"monster_name\":\"牛魔王\",\"monster_age\":500,\"Birthday\":\"2011-11-11\",\"Sal\":8000,\"Skill\":\"牛魔拳\"}"
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil{
		fmt.Printf("反序列化失败:%v\n", err)
	}
	fmt.Printf("反序列化后monster=%v monster.Name=%v\n", monster, monster.Name)
}

//将map进行反序列化
func unmarshalMap(){
	str := "{\"address\":\"洪崖洞\",\"age\":30,\"name\":\"红孩儿\"}"
	//定义一个map
	var a map[string]interface{}
	//注意:反序列化map,不需要make,因为make操作被封装到Unmarshal函数
	err := json.Unmarshal([]byte(str), &a)
	if err != nil{
		fmt.Printf("反序列化失败:%v\n", err)
	}
	fmt.Printf("反序列化后a=%v\n", a)
}

//演示对切片进行反序列化,我们这个切片[]map[string]interface{}
func unmarshalSlice(){
	str := "[{\"address\":\"北京\",\"age\":\"7\",\"name\":\"jack\"}," +
			"{\"address\":[\"墨西哥\",\"夏威夷\"],\"age\":\"20\",\"name\":\"tom\"}]"
	var slice []map[string]interface{}
	//反序列化，不需要make,因为make操作被封装到Unmarshal函数
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil{
		fmt.Printf("反序列化失败:%v\n", err)
	}
	fmt.Printf("反序列化后a=%v\n", slice)
}

func main(){
	unmarshalStruct()
	unmarshalMap()
	unmarshalSlice()
}
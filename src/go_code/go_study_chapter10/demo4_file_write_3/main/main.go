package main

import (
	"fmt"
	"io/ioutil"
)

func main(){
	//将d:/abc.txt文件内容导入到e:/kkk.txt
	//1.首先将d:/abc.txt 内容读取到内存
	//2.将读取到的内容写入e:/kkk. txt
	file1Path := "d:/abc.txt"
	file2Path := "e:/kkk.txt"
	data, err := ioutil.ReadFile(file1Path)
	if err!=nil{
		fmt.Printf("读取文件=%v\n", err)
		return
	}
	err = ioutil.WriteFile(file2Path, data, 0666)
	if err!=nil{
		fmt.Printf("写入文件=%v\n", err)
		return
	}
}
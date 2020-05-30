package main

import (
	"fmt"
	"os"
	"io"
)

//自己编写一个函数，接收两个文件路径srcFileName dstFileName
func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil{
		fmt.Printf("打开源文件错误:%v", err)
		return
	}
	defer srcFile.Close()
	//通过srcfile, 获取到Reader
	//打开dstFileName
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil{
		fmt.Printf("打开目标文件错误:%v", err)
		return
	}

	//通过dstFile,获取到Writer
	defer dstFile.Close()
	return io.Copy(dstFile, srcFile)
}

func main(){
	//将d:/flower.jpg文件拷贝到e:/abc.jpg
	//调用CopyFile完成文件拷贝
	srcFile := "E:/flower.jpg"
	dstFile := "F:/abc.jpg"
	_, err := CopyFile(dstFile, srcFile)
	if err != nil{
		fmt.Printf("复制错误:%v", err)
	}else{
		fmt.Printf("复制完成")
	}
}
package main
import (
	"fmt"
	"os"
)

func main(){
	// 打开一个文件
	// file对象 或 file指针 或 file文件句柄
	file, err := os.Open("d:/test.txt")
	if err != nil{
		fmt.Printf("打开文件出错:%v", err)
	}
	// 打开成功 我们看一下 file是什么
	fmt.Printf("file的内容:\n %v", file)
	// 关闭文件 标准方法
	err = file.Close()
	if err != nil{
		fmt.Printf("关闭文件出错:%v", err)
	}
}


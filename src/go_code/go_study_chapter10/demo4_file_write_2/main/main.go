package main

import (
	"fmt"
	"bufio"
	"os"
	"io"
)

func main(){
	filePath := "d:/abc.txt"
	// os.O_RDWR 读写方式打开
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	if err!=nil{
		fmt.Printf("open file err=%v\n", err)
		return
	}
	//及时关闭file句柄
	defer file.Close()
	reader := bufio.NewReader(file)
	// 循环读取文件的内容
	for {
		str, err := reader.ReadString('\n') // 一次读到换行结束
		if err == io.EOF{ //io.EOF 文件末尾
			break
		}
		//输出内容
		fmt.Print(str)
	}
	//准备写入5句"hello, 追加一些内容"
	str := "hello, 追加一些内容\n" // \n表示换行
	//写入时，使用带缓存的*Writer
	writer := bufio.NewWriter(file)
	for i:=0;i<5;i++{
		writer.WriteString(str)
	}
	//因为writer是带缓存，因此在调用writerstring方法时，其实
	//内容是先写入到缓存的，所以需要调用Flush方法,将缓冲的数据真正写入到文件中，
	//否则文件中会 没有数据!!! 
	writer.Flush()
}
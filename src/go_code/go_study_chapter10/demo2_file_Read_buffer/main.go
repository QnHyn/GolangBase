package main
import (
	"fmt"
	"os"
	"bufio"
	"io"
)

func main(){
	file, err := os.Open("d:/test.txt")
	if err != nil{
		fmt.Printf("打开文件出错:%v", err)
	}
	// 当函数退出时及时关闭file
	defer file.Close() // 否则会有内存泄露

	// 创建一个Reader 这里是带缓冲的 默认的缓冲大小为4096字节
	// 可以处理比较大的文件
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
	fmt.Println("文件读取结束。。。")
}


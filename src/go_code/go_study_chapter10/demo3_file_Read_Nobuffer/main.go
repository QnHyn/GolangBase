package main
import (
	"fmt"
	"io/ioutil"
)

func main(){
	//使用ioutil.ReadFile一次性将文件读取到位 返回byte切片
	file := "d:/test.txt"
	content, err := ioutil.ReadFile(file)
	if err!=nil{
		fmt.Printf("read file err=%v", err) 
	}
	//把读取到的内容显示到终端
	//fmt.Printf("%v", content) // []byte 
	fmt.Printf( "%v", string(content)) // []byte
	//我们没有显式的open文件，因此也不需要显式的close文件
	//因为，文件的Open和Close被封装到ReadFile 函数内部
}


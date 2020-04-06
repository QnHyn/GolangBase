package main
import (
	"fmt"
	"errors"
)

//函数去读取以配置文件init. conf的信息
//如果文件名传入不正确，我们就返回一个自定义的错误
/*
panic: 读取文件错误..

goroutine 1 [running]:
main.test02()
        E:/GolangBase/src/go_code/go_study_chapter07/demo13_func_error_catch/userdefined/main/main.go:23 +0xb2
main.main()
        E:/GolangBase/src/go_code/go_study_chapter07/demo13_func_error_catch/userdefined/main/main.go:30 +0x29
exit status 2
*/
func readConf(name string) (err error) {
	if name == "config.ini"{
		//读取..
		return nil
	} else {
		//这回一个目定义错误
		return errors.New("读取文件错误..")
	}
}

func test02() {
	err := readConf("config2.ini")
	if err!=nil{
		//如果读取文件发送错误，就输出这个错误，并终止程序
		panic(err)
		fmt.Println("test02()继续执行....") 
	}
}

func main(){
	//测试自定义错误的使用
	test02()
	fmt.Println("main()下面的代码...")
}
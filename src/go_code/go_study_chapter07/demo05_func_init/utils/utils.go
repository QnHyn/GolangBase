package utils 
import "fmt"
var Age int
var Name string


//Age 和Name 全局变量，我们需要在main.go 使用
//但是我们需要初始化Age和Name
//init函数完成初始化工作
func init(){
	fmt.Println("utils包的init()...")
	Age = 100
	Name = "tom~"
}

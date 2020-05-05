package main
import (
	"fmt"
	"go_code/go_study_chapter09/demo03_oop_init/model"
)


func main() {
	//如果model包的结构体变量首字母大写，引入后，直接使用,没有问题
	//var stu1 = model.Student{"小明", 19.1}
	var stu1 = model.NewStudent("小明", 19.1)
	fmt.Println(*stu1)
	fmt.Println(stu1.Name, stu1.GetScore())
}

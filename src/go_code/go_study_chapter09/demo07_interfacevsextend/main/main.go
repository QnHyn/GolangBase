package main
import (
	"fmt"
)

// Monkey 结构体
type Monkey struct{
	Name string
}

// 声明接口
type BirdAble interface{
	Flying()
}

type FishAble interface{
	Swimming()
}

func (this *Monkey) climbing(){
	fmt.Println(this.Name, "生来会爬树...")
}

// LittleMonkey机构体
type LittleMonkey struct{
	Monkey // 继承
}

func (this *LittleMonkey) Flying(){
	fmt.Println(this.Name, "通过学习会飞翔...")
}

func (this *LittleMonkey) Swimming(){
	fmt.Println(this.Name, "通过学习会游泳...")
}
func main() {
	monkey := LittleMonkey{
		Monkey{
			Name : "悟空",
		},
	}
	monkey.climbing()
	monkey.Flying()
	monkey.Swimming()
}

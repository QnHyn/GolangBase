package main
import (
	"fmt"
)

type Usb interface{
	// 声明了两个没有实现的方法
	Start()
	Stop()
}

type Phone struct{
	Name string
}

// 让结构体Phone实现Usb接口
func (p Phone) Start(){
	fmt.Println("手机开始工作。。。")
}

func (p Phone) Stop(){
	fmt.Println("手机停止工作。。。")
}

func (p Phone) Call(){
	fmt.Println("手机正在打电话")
}

type Camera struct{
	Name string
}

// 让结构体Camera实现Usb接口
func (c Camera) Start(){
	fmt.Println("相机开始工作。。。")
}

func (c Camera) Stop(){
	fmt.Println("相机停止工作。。。")
}

type Computer struct{

}

// 编写一个方法Working方法，接收一个Usb接口类型变量
// 只要是实现了Usb接口(所谓实现Usb接口，就是指实现了Usb接口声明所有方法)
func (c Computer) Working(usb Usb){
	usb.Start()
	//如果usb是指向Phone结构体变量，则还需要调用Call方法
	//类型判断.[注意体会!!]
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}
	usb.Stop()
}

func main(){
	//定义一个Usb接口数组，可以存放Phone和Camera的结构体变量
	//这里就体现出多态数组
	var usbArr [3]Usb 
	usbArr[0] = Phone{"vivo"} 
	usbArr[1]= Phone{"小米"}
	usbArr[2] = Camera{"尼康"}
	//Phone还有一个特有的方法call(), 请遍历Usb数组, 如果是Phone变量,
	//除了调用Usb接口声明的方法外, 还需要调用Phone 特有方法call.=》类型断言
	var computer Computer
	for _, v := range usbArr{
		computer.Working(v)
		fmt.Println()
	}
	fmt.Println(usbArr)
}

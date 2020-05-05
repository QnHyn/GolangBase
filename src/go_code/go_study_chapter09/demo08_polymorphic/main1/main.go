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
	usb.Stop()
}

func main(){
	//定义一个Usb接口数组，可以存放Phone和Camera的结构体变量
	//这里就体现出多态数组
	var usbArr [3]Usb 
	usbArr[0] = Phone{"vivo"} 
	usbArr[1]= Phone{"小米"}
	usbArr[2] = Camera{"尼康"}
	fmt.Println(usbArr)
}

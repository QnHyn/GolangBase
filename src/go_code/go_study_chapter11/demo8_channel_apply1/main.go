package main
import(
	"fmt"
	"time"
)

// 写数据
func writeData(intChan chan int){
	for i := 1; i <= 500; i++{
		intChan<- i
		fmt.Println("写入管道数据:", i)
	}
	close(intChan)
}
func readData(intChan chan int, exitChan chan bool){
	for {
		v, ok := <-intChan
		if !ok{
			break
		}
		// 频率不一样完全没问题
		//time.Sleep(time.Second)
		fmt.Println("读到管道数据:", v)
	}
	exitChan<- true
	close(exitChan)
}
func main(){
	// 创建两个管道
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)
	// 启动 写读协程
	go writeData(intChan)
	go readData(intChan, exitChan)
	// 不停读取退出管道, 解决之前不知道何时协程执行完成的问题
	for {
		// 直到可以读到exitChan的值
		_, ok := <-exitChan
		// fmt.Println(ok) 这里!ok和ok是一样的结果
		// 因为!ok指的取到后再取一次取不到, ok是直接取到(更直接) 
		if !ok{
			break
		}
	}
}
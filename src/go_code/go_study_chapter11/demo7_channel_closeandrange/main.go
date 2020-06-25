package main
import(
	"fmt"
)

func main(){
	intchan := make(chan int, 3)
	intchan <- 10
	intchan <- 20
	// 关闭管道
	close(intchan)
	// 关闭后不能在存放
	// intchan <- 30
	// 读取数据完全没有问题
	n1 := <-intchan
	<- intchan
	fmt.Println(n1)
	
	intchan2 := make(chan int, 100)
	for i := 0; i < 100; i++{
		intchan2<- i*2
	}
	//	遍历管道不能使用普通的for循环 因为一个取出后第二个就变成第一个位置啦
	// for i := 0; i < len(intChan2); i++ {
	// }
	// 遍历前需要先关闭管道
	close(intchan2)
	for v := range intchan2{
		fmt.Println(v)
	}
}
package main
import(
	"fmt"
	"time"
)

func putNum(intChan chan int){
	for i := 1; i <= 20000; i++{
		intChan<- i
	}
	close(intChan)
}
// 取值判断是否为素数
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool){
	var flag bool
	for{
		//time.Sleep(time.Millisecond * 10)
		num, ok :=<- intChan
		if !ok{
			break
		}
		flag = true
		// 判断是不是素数
		for i := 2; i < num; i++{
			if num % i == 0{
				flag = false
				break
			}
		}
		if flag{
			primeChan<- num
		}
	}
	fmt.Printf("有协程已经取不到数据啦, 退出")
	exitChan<- true
}
func main(){
	intChan := make(chan int, 20000)
	primeChan := make(chan int, 10000)
	exitChan := make(chan bool, 4)

	// 统计用时
	start := time.Now().Unix()
	// 开启一个协程放入1-8000个数
	go putNum(intChan)
	// 开启四个协程取出数据判断素数
	for i := 0; i < 4; i++{
		go primeNum(intChan, primeChan, exitChan)
	}
	// 再开一个协程 等待标志管道取出四个值
	go func(){
		for i := 0; i < 4; i++{
			<- exitChan
		}
		// 取出后关闭管道primeChan
		close(primeChan)
	}()
	// 遍历我们的primeChan把结果取出
	for {
		res, ok := <-primeChan
		if !ok{
			break
		}
		fmt.Printf("素数%d\n", res)
	}
	end := time.Now().Unix()
	fmt.Printf("用时:%v", end - start)
}
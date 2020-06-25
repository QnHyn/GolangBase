package main
import(
	"fmt"
	"time"
	"sync"
)

// 定义一个全局map
var(
	myMap = make(map[int]int, 10)
	// 声明一个全局的互斥锁 lock是一个全局的互斥锁 
	lock sync.Mutex
)

// 计算n!结果以及之前结果保存在myMap中
func test(n int){
	res := 1
	for i := 1; i <= n; i++{
		res *= i
	}
	// 将结果储存在myMap中
	// 加锁
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}

func main(){
	for i := 1; i < 20; i++{
		go test(i)
	}

	// 休眠10秒 防止主线程退出 其他协程也退出
	time.Sleep(time.Second*10)
	// 输出变量结果
	// 主线程并不知道10秒能执行完成，因此底层可能仍然出现资源争夺，因此加入互斥锁即可解决问题
	lock.Lock()
	for i, v := range myMap{
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	lock.Unlock()
}
package main
import(
	"fmt"
	"time"
)

// 定义一个全局map
var(
	myMap = make(map[int]int, 10)
)

// 计算n!结果以及之前结果保存在myMap中
func test(n int){
	res := 1
	for i := 1; i <= n; i++{
		res *= i
	}
	// 将结果储存在myMap中
	myMap[n] = res
}

func main(){
	for i := 1; i < 20; i++{
		go test(i)
	}

	// 休眠10秒 防止主线程退出 其他协程也退出
	time.Sleep(time.Second*10)
	// 输出变量结果
	for i, v := range myMap{
		fmt.Printf("map[%d]=%d\n", i, v)
	}
}
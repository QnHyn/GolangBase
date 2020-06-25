package main
import(
	"fmt"
	"runtime"
)

func main(){
	//获取当前系统逻辑cpu的数量
	num := runtime.NumCPU()
	//GOMAXPROCS设置可同时执行的最大CPU数，并返回先前的设置
	runtime.GOMAXPROCS(num)
	fmt.Println("num=", num)
}
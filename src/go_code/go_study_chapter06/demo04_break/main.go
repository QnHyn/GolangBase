package main
import (
	"fmt"
	"math/rand"
	"time"
) 


func main() {
	// 为了生成一个随机数,还需要给rand设置一个种子(否则是一个伪随机数一直是82)
	//time .Now().Unix() :返回一个从1970:01:01的0时0分0秒到现在的秒数
	rand.Seed(time.Now().Unix())
	fmt.Println(time.Now().Unix())
	//随机生成一个1-100的数rand.Intn(100)[0, 100)
	n := rand.Intn(100) + 1
	fmt.Println(n)

	// 计数i
	var i int = 0
	for {
		rand.Seed(time.Now().Unix())
		k := rand.Intn(100) + 1
		i++
		if k == 99{
			fmt.Println(i)
			break
		}
	}

	// 指定标签形式 跳转到指定break
	label1:
	for i := 0; i < 4; i++ {
		//label2:
		for j := 0; j < 10; j++ {
			if j == 2{ 
				//break // 默认只会跳出最近的for循环
				break label1
			}
			fmt.Println("j的值是:", j)
		}
	}
}
package main
import (
	"fmt"
	"math/rand"
	"time"
)


func main() {
	//创建一个byte类型的26个元素的数组，分别放置'A'-'Z'。
	//使用for循环访问所有元素并打印出来。提示:字符数据运算'A'+1 -> 'B'
	var byteArr [26]byte
	for i := 0; i < len(byteArr); i++{
		byteArr[i] = byte('A' + i)
		fmt.Printf("byteArr[%v]的值为%c\n", i, byteArr[i])
	}

	//请求出一个数组的最大值，并得到对应的下标
	//思路
	//1. 声明一个数组var intArr = [...]int{1, -1, 9, 90, 11}
	//2.假定第一个元素就是最大值，下标就e
	//3.然后从第二个元素开始循环比较，如果发现有更大，则交换
	var intArr = [...]int{1, -1, 9, 90, 11}
	max_val := intArr[0]
	for i := 1; i< len(intArr); i++{
		if max_val < intArr[i]{
			max_val = intArr[i]
		}
	}
	fmt.Println("最大值为:", max_val)


	//请求出一个数组的和和平均值。for-range
	//1.就是声明一个数组var intArr= [...]int{1, -1, 9, 90，11} 
	//2.求出和sum
	//3.求出平均值
	var sum int
	for _, v := range(intArr){
		sum += v
	} 
	avr := float64(sum) / float64(len(intArr))
	fmt.Println(sum, avr)

	
	//要求:随机生成五个数，并将其反转打印
	//思路
	//1.随机生成五个数, rand.Intn()函数
	//2.当我们得到随机数后，就放到一个数组int数组
	//3. 反转打印，交换的次数是len/2,倒数第一个和第一个元素交换，倒数第2个和第2个元素交
	var intArr3 [5]int
	//为了每次生成的随机数不一样，我们需要给个seed值
	len := len(intArr3)

	// Go的rand包会使用相同的源来产生一个确定的伪随机数序列。这个源会产生一个不变的数列
	rand.Seed(time.Now().UnixNano())
	for i:=0; i<len; i++{
		intArr3[i] = rand.Intn(100) // 0<=n<100 
	}
	fmt.Println("交换前~=", intArr3)
	//反转打印,交换的次数是 len / 2,
	//倒数第一个和第一个元素交换，倒数第2个和第2个元素交换
	temp := 0 //做一个临时变量
	for i:=0; i<len/2; i++{
		temp = intArr3[len - 1 - i]
		intArr3[len - 1 - i] = intArr3[i]
		intArr3[i] = temp 
	}
	fmt.Println("交换后~=", intArr3)
}

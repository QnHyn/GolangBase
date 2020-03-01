package main
import "fmt"


func main() {
	for i := 1; i <= 10; i++{
		fmt.Println("for 循环的基础用法", i)
	}

	// for循环第二种使用方式
	j := 1 // 循环变量的初始化
	for j < 10 { // 循环条件
		fmt.Println("for 循环的第二种使用方法", j)
		j++ //循环变量迭代
	}

	// for 循环第三种使用方法
	k := 1
	for{
		fmt.Println("for 循环的第三种使用方法")
		k++
		if k > 3 {
			break
		}			
	}

	// for-range 方便遍历字符串或者数组
	var str string = "hello for使用"
	
	// 传统方式 需要用切片解决中文问题
	str2 := []rune(str) //就是把str转成[]rune
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c", str2[i])
	}

	// for-range方式遍历
	for index, val := range str{
		//如果我们的字符串含有中文，
		//那么传统的遍历字符串方式，就是错误，会出现乱码。
		//原因是传统的对字符串的遍历是按照字节来遍历。
		//for-range方式不会出错
		fmt.Printf("index=%d,value=%c \n", index, val)
	}
}
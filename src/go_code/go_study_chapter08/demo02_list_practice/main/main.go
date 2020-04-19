package main
import "fmt"


/*
	[0 0 0]
	[10 20 30]
	intArr的地址=0xc0420520a0 intArr[0] 地址0xc0420520a0 intArr[1] 地址0xc0420520a8 intArr[2] 地址0xc0420520b0PS
*/
func main() {
	var intArr [3]int //int占8个字节
	//当我们定义完数组后，其实数组的各个元素有默认值0
	fmt.Println(intArr)
	intArr[0] = 10 
	intArr[1] = 20
	intArr[2] = 30
	fmt.Println(intArr)
	fmt.Printf("intArr的地址=%p intArr[0] 地址%p intArr[1] 地址%p intArr[2] 地址%p",
	&intArr, &intArr[0], &intArr[1], &intArr[2])


	//四种初始化数组的方式
	var numArr01 [3]int = [3]int{1, 2, 3}
	fmt.Println("numArr01=", numArr01)
	var numArr02 = [3]int{5, 6, 7}
	fmt.Println("numArr02=", numArr02)
	//这里的 [...]是规定的写法
	var numArr03 = [...]int{8, 9, 10}
	fmt.Println("numArr03=", numArr03)
	var numArr04 = [...]int{1: 800, 0: 900, 2:999}
	fmt.Println("numArr04=", numArr04)
	//类型推导
	strArr05 := [...]string{1: "tom", 0: "jack", 2:"mary"}
	fmt.Println("strArr05=", strArr05)


	//演示for-range遍历数组
	heroes := [...]string{"宋江", "吴用", "卢俊义"}
	//使用常规的方式遍历
	for i := 0; i < len(heroes); i++{
		fmt.Printf("常规方式遍历heroes[%d]=%v\n", i, heroes[i])
	}
	//for-range遍历数组
	for i, v := range heroes {
		fmt.Printf("i=%v v=%v\n", i, v)
		fmt.Printf("for-range遍历heroes[%d]=%v\n", i, heroes[i])
	}
	for _, v := range heroes {
		fmt.Printf("for-range遍历元素的值=%v\n", v)
	}

	arr := [3]int{11, 22, 33}
	test(arr)
	fmt.Println(arr)
	test01(&arr)
	fmt.Println(arr)

}
// 直接传递数组 值传递 复制一份数组
func test(arr [3]int) {
	arr[0] = 88
}
// 把数组的地址传过来
func test01(arr *[3]int) {
	(*arr)[0] = 88
}
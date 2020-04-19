package main
import "fmt"
	

func main() {
	var arr [5]int = [...]int{1,2,3,4,5}
	// 第一种方式: 定义一个切片，然后让切片去引用一个已经创建好的数组
	var slice1 = arr[1:3]
	fmt.Println("arr=", arr)
	fmt.Println("slice1=", slice1) 
	fmt.Println("slice1 len =", len(slice1))
	fmt.Println("slice1 cap =", cap(slice1))

	//第二种方式: 通过make来创建切片
	var slice2 []float64 = make([]float64, 5, 10)
	slice2[1] = 10
	slice2[3] = 20
	//对于切片，必须make使用。
	fmt.Println(slice2)
	fmt.Println("slice2的size=", len(slice2))
	fmt.Println("slice2的cap=", cap(slice2))

	
	//第3种方式:定义一个切片，直接就指定具体数组，使用原理类似make的方式
	var strslice []string = []string{"tom", "jack" , "mary"}
	fmt.Println("strslice=", strslice)
	fmt.Println("strslice size=", len(strslice)) //3
	fmt.Println("strslice cap=", cap(strslice)) // 3


	//使用常规的for循环遍历切片
	var arrint [5]int = [...]int{10, 20, 30, 40, 50} 
	sliceint := arrint[1:4] // 20,30,40
	for i:=0; i<len(sliceint); i++{
		fmt.Printf("sliceint[%v]=%v", i, sliceint [i]) 
	}
	fmt.Println()
	//使用for-range方式遍历切片
	for i, v := range sliceint  {
		fmt.Printf("i=%v v=%v \n", i, v)
	}
}

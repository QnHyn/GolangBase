package main
import "fmt"
	

func main() {
	//使用常规的for循环遍历切片
	var arr [5]int = [...]int{10, 20, 30, 40, 50}
	slice := arr[1:4]
	// 对切片进行切片
	slice2 := slice[1:2] // slice[20, 30, 40]
	slice2[0] = 100
	//因为arr，slice 和slice2 指向的数据空间是同一个，因此slice2[0]=100,
	fmt.Println("slice2=", slice2)
	fmt.Println("slice=", slice)
	fmt.Println("arr=", arr)

	//用append内置函数，可以对切片进行动态追加
	var slice3 []int = []int{100, 200, 300}
	//通过append直接给slice3追加具体的元素
	slice3 = append(slice3, 400, 500, 600)
	fmt.Println("slice3", slice3) //100, 200, 300, 400, 500, 600
	//通过append将切片slice3追加给slice3 把slice追加给slice3 都要加 ...
	// 不能直接追加数组slice3 = append(slice3, arr)
	slice3 = append(slice3, arr[1:2]...)
	slice3 = append(slice3, slice...)
	slice3 = append(slice3, slice3...) // 100, 200, 300,400, 500, 600 100, 200, 300, 4
	fmt.Println("slice3", slice3)


	//切片的拷贝操作
	//切片使用copy内置函数完成拷贝，举例说明
	fmt.Println()
	var slice4 []int = []int{1, 2, 3, 4, 5}
	var slice5 = make([]int, 10)
	copy(slice5, slice4)
	fmt.Println(" slice4=", slice4)// 1, 2, 3, 4, 5
	fmt.Println("slice5=", slice5) // 1, 2, 3, 4, 5, 0，0 ,0,0,0

	// a的长度大于slicea 把a复制过去不会报错
	var a []int = []int {1,2,3,4,5} 
	var slicea = make([]int, 1)
	fmt.Println(slicea) // [0]
	copy(slicea, a)
	fmt.Println(slicea)// [1]
}

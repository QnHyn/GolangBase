package main
import(
	"fmt"
)

type Cat struct{
	Name string
	Age int
}

func main(){
	// 创建一个map管道
	var mapChan chan map[string]string
	mapChan = make(chan map[string]string, 5)
	m1 := make(map[string]string, 10)
	m2 := make(map[string]string, 20)
	m1["city1"] = "北京"
	m1["city2"] = "上海"
	m2["hero1"] = "松江"
	m2["hero2"] = "红孩儿"
	mapChan<- m1
	mapChan<- m2
	//可以直接从管道中扔出
	<- mapChan
	<- mapChan
	fmt.Printf("Map管道%v\nm1的内容%v\nm2的内容%v\n", mapChan, m1, m2)

	// 创建一个catChan,最多可以存放10个Cat结构体变量,演示写入和读取
	var catChan chan Cat
	catChan = make(chan Cat, 10)
	cat1 := Cat{Name:"Tom", Age:18}
	cat2 := Cat{Name:"Tom~", Age:23}
	catChan<- cat1
	catChan<- cat2
	cat11 := <-catChan
	cat22 := <-catChan
	fmt.Printf("cat11的内容%v\ncat22的内容%v\n", cat11, cat22)

	// 创建一个catChan2, 最多可以存放10个*Cat变量,演示写入和读取的用法
	var catChan2 chan *Cat
	catChan2 = make(chan *Cat, 10)
	catChan2<- &cat1
	catChan2<- &cat2
	cat_1 := <-catChan2
	cat_2 := <-catChan2
	fmt.Println(cat_1, cat_2)

	// 创建一个allChan,最多可以存放10个任意数据类型变量,演示写入和读取的用法
	var allChan chan interface{}
	allChan = make(chan interface{}, 10)
	allChan<- cat1
	allChan<- cat2
	allChan<- 10
	allChan<- "jack"
	cat1_all := <- allChan
	cat2_all := <- allChan
	num_all := <-allChan
	str_all := <-allChan
	fmt.Println(cat1_all, cat2_all, num_all, str_all)
	// 注意下面会报错 interface {} is interface with no methods
	// 在编译层面它认为cat2_all是空接口类型
	//fmt.Println(cat2_all.Name)
	// 可以用类型断言
	a := cat2_all.(Cat)
	fmt.Println(a.Name)
}
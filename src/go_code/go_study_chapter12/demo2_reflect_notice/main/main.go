package main
import (
	"fmt"
	"reflect"
)

func testInt(b interface{}) { 
	val := reflect.ValueOf(b)
	fmt.Printf("val type=%T\n", val)
	// SetInt不能用于地址 所以需要用到Elem的函数 进行指向
	val.Elem().SetInt(110)
	fmt.Printf("val=%v\n", val)
}

func main(){
	var num int = 20
	testInt(&num)
	fmt.Println("num=", num)

	//你可以这样理解rVal.Elem( )
	// num := 9
	// ptr *int = &num
	// num2 := *ptr //类似rVal.Elem()
}
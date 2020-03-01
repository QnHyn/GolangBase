package main
import "fmt"


func main() {
	var age int = 20
	// switch 不加表达式 当作if-else使用
	switch {
		case age == 10:
			fmt.Println("age的值为10")
		case age >= 20:
			fmt.Println("age的值为20以上")
		default:
			fmt.Println("默认")
	}

	// switch 后也可以直接声明定义一个变量，分号结束，不推荐
	switch score := 80; {
	case score < 60:
		fmt.Println("不合格")
	case score > 60 && score <=80:
		fmt.Println("合格")
	default:
		fmt.Println("优秀")
	}

	// fallthrough switch穿透
	var num int = 20
	switch num{
		case 20:
			fmt.Println("num的值为20")
			fallthrough //默认只能穿透一层 输出为num的值为20 num的值为30
		case 30:
			fmt.Println("num的值为30")
		default:
			fmt.Println("默认")
	}

	//switch 语句还可以被用于type-switch来判断某个interface(接口)变量中实际指向的变量类型
	var x interface{}
	var y = 10.0
	x = y
	switch i := x.(type) {
	case nil:
		fmt.Printf(" x的类型~ :%T" ,i)
	case int:
		fmt. Printf("x是int型")
	case float64:
		fmt. Printf("x是float64型" )
	case func(int) float64:
		fmt.Printf("x是func(int) 型")
	case bool, string:
		fmt.Printf("x是bool或string 型")
	default:
		fmt.Printf("未知型")
	}
}
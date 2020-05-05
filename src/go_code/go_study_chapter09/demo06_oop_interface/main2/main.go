package main
import (
	"fmt"
)

type AInterface interface{
	Say()
}
type BInterface interface{
	Hello()
}

type Monster struct{

}

func (m Monster) Say(){
	fmt.Println("Monster say~~")
}
func (m Monster) Hello(){
	fmt.Println("Monster hello~~")
}
type T interface{

}
type Stu struct{
	Name string
}
func main() {
	var m Monster
	var a AInterface = m
	var b BInterface = m
	a.Say()
	b.Hello()

	var stu Stu
	var t T = stu//ok
	fmt.Println(t)
	var t2 interface{} = stu
	var num1 float64 = 8.8 
	t2 = num1
	t = num1
	fmt.Println(t2, t)
}

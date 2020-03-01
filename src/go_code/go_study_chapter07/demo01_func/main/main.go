package main
import (
	"fmt"
	"go_code/go_study_chapter07/demo01_func/utils"
)


func main() {
	var n1 float64 = 4.5
	var n2 float64 =  1.5
	var operater byte = '+'
	res := utils.Cal(n1, n2, operater)
	fmt.Println("res:", res)
}
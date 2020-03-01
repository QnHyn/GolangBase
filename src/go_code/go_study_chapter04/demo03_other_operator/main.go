package main
import "fmt"


func main() {
	var a int = 10
	var b int = 9

	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("a的值%d, b的值%d \n", a, b)
}
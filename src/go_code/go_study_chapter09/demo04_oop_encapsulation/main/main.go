package main
import (
	"fmt"
	"go_code/go_study_chapter09/demo04_oop_encapsulation/model"
)


func main() {
	var p1 = model.NewPerson("小明")
	p1.SetAge(10)
	p1.SetSalary(10000.0)
	fmt.Println(p1.Name, p1.GetAge(), p1.GetSalary())
}

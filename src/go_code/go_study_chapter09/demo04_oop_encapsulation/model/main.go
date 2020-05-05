package model
import (
	"fmt"
)


type person struct{
	Name string
	age int // 小写其他包不能直接访问
	sal float64
}


func NewPerson(name string) *person{
	return &person{
		Name : name,
	}
}

func (p *person) SetAge(age int) {
	if age > 0 && age < 150{
		p.age = age
	}else{
		fmt.Println("输入年龄错误！")
	}
}

func (p *person) GetAge() int{
	return p.age
}


func (p *person) SetSalary(sal int) {
	if sal >= 3000 && sal <= 30000{
		p.age = sal
	}else{
		fmt.Println("薪水输入错误！")
	}
}

func (p *person) GetSalary() float64{
	return p.sal
}
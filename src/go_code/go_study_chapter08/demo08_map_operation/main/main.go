package main
import "fmt"
	

func main() {
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"
	fmt.Println(cities)
	//因为no3这个key已经存在，因此下面的这句话就是修改
	cities["no3"]="上海~"
	fmt.Println(cities)

	//演示删除
	delete(cities, "no1" )
	fmt.Println(cities)
	//当delete指定的key不存在时，删除不会操作，也不会报错
	delete(cities, "no4")
	fmt.Println(cities)

	//如果希望一次性删除所有的key
	//1.遍历所有的key,如何逐一删除[遍历]
	//2.直接make一个新的空间
	cities = make(map[string]string)
	fmt.Println(cities)

	//演示map的查找
	val, ok := cities["no2"]
	if ok{
		fmt.Printf("有no1 key值为%v\n", val)
	} else {
		fmt.Printf("没有no1 key\n") 
	}

	studentMap := make(map[string]map[string]string)
	studentMap["stu01"] = make(map[string]string, 3)
	studentMap["stu01"]["name"] = "tom"
	studentMap["stu01"]["sex"] = "男"
	studentMap["stu01"]["address"] = "北京长安街~"

	studentMap["stu02"] = make(map[string]string, 3) //这句话不能少!!
	studentMap["stu02"]["name"] ="mary"
	studentMap["stu02"]["sex"] = "女"
	studentMap["stu02"]["address"] = "上海黄浦江~"
	
	// map的遍历
	for k1, v1 := range studentMap {
		fmt.Println("k1=", k1)
		for k2, v2 := range v1 {
			fmt.Printf("\t k2=%v v2=%v\n", k2, v2)
			fmt.Println()
		}
	}

	fmt.Println(len(studentMap))
}


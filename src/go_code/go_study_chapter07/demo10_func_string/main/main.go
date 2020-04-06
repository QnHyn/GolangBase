package main
import (
	"fmt"
	"strconv"
	"strings"
)


func main(){
	var str string = "hello"
	fmt.Println("str的长度:", len(str)) //5
	
	// 如果字符串中有中文 golang的编码统一为utf-8(字符数字ASCII中占1个字节, 汉字占三个字节)
	var str1 string = "hello北"
	fmt.Println("str1的长度:", len(str1)) //8
	for i := 0; i < len(str1); i++{
		fmt.Printf("字符=%c\n", str1[i]) // 中文部分会乱码
	}

	// 通过r := []rune(str)解决中文的问题
	r := []rune(str1)
	for i := 0; i < len(r); i++{
		fmt.Printf("字符=%c\n", r[i]) // 中文部分会乱码
	}

	// 字符串转整数 n, err := strconv.Atoi("12") 可以进行数据的校验
	n, err := strconv.Atoi("12")
	//n, err := strconv.Atoi("hello")
	if err != nil {
		fmt.Println("转换错误", err)
	}else {
		fmt.Println("转成的结果是", n)
	}

	// 整数转字符串
	str = strconv.Itoa(12345)
	fmt.Printf("str转成的结果是%v, \nstr类型是%T", str, str)

	// 字符串转[]byte切片 var bytes = []byte("hello go")
	var bytes = []byte("he1lo go") 
	fmt.Println("bytes=%v", bytes)

	// []byte 转字符串: str = string([]byte{97, 98, 99})
	str = string([]byte{97, 98, 99})
	fmt.Println(str)

	//10进制转2, 8，16进制: str = strconv.FormatInt(123, 2), 返回对应的字符串
	str = strconv.FormatInt(123, 2)
	fmt.Printf("123对应的二进制是=%v\n", str)
	str = strconv.FormatInt(123, 16)
	fmt.Printf("123对应的16进制是=%v\n", str)

	//查找子串是否在指定的字符串中: strings.Contains("seafood", "foo") //true
	b := strings.Contains("seafood", "mary")
	fmt.Printf("b=%v\n", b)

	//统计一个字符串有几个指定的子串: strings.Count("ceheese", "e") //4
	num := strings.Count("ceheese", "e" )
	fmt.Printf("num=%v\n", num)

	//返回子串在字符串第一次出现的index值，如果没有返回-1:strings.Index("NLT_ abc", "abc") // 4
	index := strings.Index("NLT_abcabcabc", "abc") // 4
	fmt.Printf("index=%v\n", index)
	
	//返回子串 在字符串最后一次出现的index,如没有返回-1 : strings ,LastIndex("go golang", "go")
	index = strings.LastIndex("go golang", "go") //3
	fmt.Printf("index=%v\n", index)

	//将指定的子串替换成另外一个子串: strings. Replace("go go hello", "go", "go语言"，n)
	//n可以指定你希望替换几个，如果n=-1表示全部替换
	str2 := "go go hello"
	str = strings.Replace(str2, "go", "北京", -1)
	fmt.Printf("str=%v str2=%v\n", str, str2)

	//按照指定的某个字符，为分割标识，将一个字符串拆分成字符串数组:
	//strings.Split("hello,wrold,ok", ”,”)
	strArr := strings.Split("hello,wrold,ok" , "," )
	for i := 0; i < len(strArr); i++{
		fmt.Printf("str[%v]=%v\n", i, strArr[i])
		fmt.Printf("strArr=%v\n", strArr)
	}

	//将字符串的字母进行大小写的转换:strings.ToLower("Go") // go strings.ToUpper("Go")
	str = "goLang Hello"
	str = strings.ToLower(str)
	str = strings.ToUpper(str)
	fmt.Printf("str=%v\n", str) //golang hello

	//将字符串左右两边的空格去掉: strings.Trimspace(" tn a lone gopher ntrn")
	str = strings.TrimSpace(" tn a lone gopher ntrn")
	fmt. Printf("str=%q\n", str)

	//将字符串左右两边指定的字符去掉:strings.Trim("! hello!","!") // ["hello"] //将左右两边!和”"去掉
	str = strings.Trim("!hello!","!")
	fmt.Printf("str=%q\n", str)

	//判断字符串是否以指定的字符串开头:
	//strings .HasPrefix("ftp://192.168.10.1". "ftp") // true 
	b = strings.HasPrefix("ftp://192.168.10.1", "hsp") //true
	fmt.Printf("b=%v\n", b)
}
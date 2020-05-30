package main

import (
	"testing"
)

// cal_test.go cal.go中有addUpper函数
// 编写测试用例, 测试addUpper是否正确
func TestAddUpper(t *testing.T){
	res := addUpper(10) // 1.+ 10 = 55
	if res != 55{
		t.Fatalf("addupper错误返回值=%v期望值=%v\n", res, 55)
	}
	t.Logf("addupper正确返回值=%v期望值=%v\n", res, 55)
}
package TypeTest

import "testing"

type MyInt int64

/**
与其他主要编程语⾔的差异:
	1. Go 语⾔不允许隐式类型转换
	2. 别名和原有类型也不能进⾏隐式类型转换

类型的预定义值
	1. math.MaxInt64
	2. math.MaxFloat64
	3. math.MaxUint32
*/
func TestType(t *testing.T) {
	var a int64 = 1
	var b int = 2
	var c int16 = 3
	var d MyInt = 4
	a = int64(b) //只能强制类型转换
	//a = b //报错提示不能将int32类型隐式转化成为int64类型
	//a = d//同样报错，即便他们是同类型但是不同名也不支持隐式类型转换
	t.Log(a, b, c, d)
}

/**
1. 不⽀持指针运算
2. string 是值类型，其默认的初始化值为空字符串，⽽不是 nil
*/
func TestPtr(t *testing.T) {
	a := 1
	aPtr := &a
	//aPtr += 1//编译错误！指针不支持运算
	t.Log(a, aPtr)
	t.Logf("%T,%T", a, aPtr) //%T输出变量类型
}
func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
	if s == "" {
		t.Log("是空串")
	}
	//if s == nil {//s会被初始化为空串而不是空
	//	t.Log("是空")
	//}
}

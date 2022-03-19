package fib

import "testing"

/**
1. 首先在测试文件中我们的方法名以Test开头，形式参数是t *tesing t
2. go语言中利用var关键词来进行赋值，var a int 的格式---变量类型在变量名后面，var具有一定的类型推断能力，如果用的是默认类型我们可以省略类型。
	另外也可以用var加上圆括号的形式来进行一次性多个变量声明赋值
	还可以直接使用类型推断运算符=:直接进行声明赋值，a1:=1
3.在测试文件里面我们可以直接不使用fmt.Print来打印，我们可以直接t.Log()

与其他主要编程语⾔的差异：
1. 赋值可以进⾏⾃动类型推断
2. 在⼀个赋值语句中可以对多个变量进⾏同时赋值，比如a,b=b,a	这一段代码可以完成两个变量值的交换
*/
func TestFib(t *testing.T) {
	//var a1 = 1
	//var a2 = 1
	//var a1 int = 1
	//var a2 int = 1
	//var (
	//	a1 = 1
	//	a2 = 1
	//)
	a1 := 1
	a2 := 1
	t.Log(a1)

	for i := 0; i < 5; i++ { //这里打印前5个斐波那契数列
		t.Log(" ", a2)
		temp := a2
		a2 += a1
		a1 = temp
	}
}
func TestExchange(t *testing.T) {
	a := 1
	b := 2
	temp := b
	b = a
	a = temp
	//a, b = b, a
	t.Log(a, b)
}

package Condition

import (
	"fmt"
	"runtime"
	"testing"
)

/**
1. condition 表达式结果必须为布尔值
2. ⽀持变量赋值：
if  var declaration;  condition {
	// code to be executed if condition is true
}
*/
func TestIfMultiSec(t *testing.T) {
	if a := 1 == 1; a {
		t.Log("1==1")
	}
	//GO语言的方法支持多返回值，这里if的赋值就很有用。如果没有错误，也就是err为空，那么就执行某一个语句，否则就执行另外一个语句
	//if a,err:=somefuc();err==nil{
	//	t.Log(a)
	//}else {
	//	t.Log("eer")
	//}
}

/**
与其他主要编程语⾔的差异:
	1. 条件表达式不限制为常量或者整数；//也可以为字符串类型
	2. 单个 case 中，可以出现多个结果选项, 使⽤逗号分隔；
	3. 与 C 语⾔等规则相反，Go 语⾔不需要⽤break来明确退出⼀个 case；
	4. 可以不设定 switch 之后的条件表达式，在此种情况下，整个 switch 结构与多个 if…else… 的逻辑作⽤等同
*/
func TestSwitchCaseMulti(t *testing.T) {
	for a := 0; a < 5; a++ {
		switch a {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("out of range!")
		}
	}
}
func TestSwitchCaseCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch { //这种格式的Switch语句switch后面不跟变量,等价于else if
		case i%2 == 0:
			t.Log("Even")
		case i%2 == 1:
			t.Log("Odd")
		default:
			t.Log("unknown")
		}
	}
}

/**
这switch前面和if很像，都可以在前面有一个赋值语句
*/
func TestSwitch_1(t *testing.T) {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	//break
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}

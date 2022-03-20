package loop

import (
	"testing"
)

/**
与其他主要编程语⾔的差异:
Go 语⾔仅⽀持循环关键字 for，没有while！！！
*/
func TestWhileLoop(t *testing.T) {
	//下面这段代码等价于其它语言while(n<5){...;n++}
	n := 0
	for n < 5 {
		t.Log(n)
		n++
	}

	//⽆限循环
	//while (true)
	//n := 0
	//for {
	//	…
	//}
}

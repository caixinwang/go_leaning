package class2

import "testing"

/**
1. 类似于赋值，我们可以一个一个赋值常量，也可以一次性多个赋值常量。另外我们可以使用iota关键词来进行连续常量值的赋值
2. 这里的iota关键词可以进行位运算的连续赋值，0001,0010,0100,1000
*/
func TestConst_1(t *testing.T) {
	const a = 2
	const (
		a1 = iota + 1
		a2
		a3
		a4
		a5
		a6
		a7
	)
	const (
		b1 = 1
		b2 = 2
		b3 = 3
		b4 = 4
		b5 = 5
		b6 = 6
		b7 = 7
	)
	const ( //iota要在后面，刚开始iota是0，表示1没有左移
		c1 = 1 << iota //0001
		c2             //0010
		c3             //0100
		c4             //1000

	)
	p := 5 //0101
	t.Log(a1, a2, a3, a4, a5, a6, a7)
	t.Log(p&c4 == c4, p&c3 == c3, p&c2 == c2, p&c1 == c1)

}

package Operator_test

import "testing"

/**


• 相同维数且含有相同个数元素的数组才可以⽐较
• 每个元素都相同的才相等

与其它语言不同：
	1. go语言的数组是值类型，可以直接进行比较
	2. Go 语⾔没有前置的 ++，--
*/
func TestArr(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 5}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	//t.Log(a == c) //报错，长度不相同的数组不能进行比较
	t.Log(a == d)

}

func TestBitClear(t *testing.T) {
	const ( //iota要在后面，刚开始iota是0，表示1没有左移
		c1 = 1 << iota //0001
		c2             //0010
		c3             //0100
		c4             //1000
	)
	p := 7 //0111
	t.Log(p&c4 == c4, p&c3 == c3, p&c2 == c2, p&c1 == c1)
	p &^= c1 //等价于p=p&^c1----把c1功能抹掉
	t.Log(p&c4 == c4, p&c3 == c3, p&c2 == c2, p&c1 == c1)

}

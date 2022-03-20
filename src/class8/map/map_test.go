package _map

import "testing"

/**
Map 的 value 可以是⼀个⽅法
与 Go 的 Dock type 接⼝⽅式⼀起，可以⽅便的实现单⼀⽅法对象的⼯⼚模式
*/
func TestMapWithFunValue(t *testing.T) {
	map1 := map[int]func(x int) int{}
	map1[1] = func(x int) int {
		return x
	}
	map1[2] = func(x int) int {
		return x * x
	}
	map1[3] = func(x int) int {
		return x * x * x
	}
	t.Log(map1[1](5), map1[2](9), map1[3](4))
}

/**
Go 的内置集合中没有 Set 实现， 可以 map[type]bool
1. 元素的唯⼀性
2. 基本操作:添加元素/判断元素是否存在/删除元素/元素个数
*/
func TestMapForSet(t *testing.T) {
	map1 := map[string]bool{}
	//加入元素
	map1["wangcaixin"] = true //加入集合的操作---添加
	map1["Jone"] = true       //加入集合的操作---添加
	/**
	判断是否存在
	*/
	n := "Jone"
	if map1[n] {
		t.Logf("%s exists", n)
	} else {
		t.Logf("%s is not existd", n)
	}
	//删除元素
	delete(map1, "wangcaixin")
	//元素个数
	t.Log(len(map1))
}

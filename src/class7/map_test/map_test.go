package map_test

import (
	"testing"
)

/**
与其他主要编程语⾔的差异:
	在访问的 Key 不存在时，仍会返回零值，不能通过返回 nil 来判断元素是否存在
*/
func TestInitMap(t *testing.T) {
	map1 := map[int]int{}                   //初始化1
	map2 := map[int]int{1: 2, 3: 12, 6: 88} //初始化2
	map3 := make(map[int]int, 10)           //初始化3
	var map4 map[int]int = map[int]int{}    //初始化4
	t.Log(map1, map2, map3, map4)
	t.Log(len(map1), len(map2), len(map3), len(map4))
}
func TestAccessNotExistingKey(t *testing.T) {
	map1 := map[int]int{}                   //初始化1
	map2 := map[int]int{1: 0, 3: 12, 6: 88} //初始化2
	/**
	两个都是0，但是一个是未初始化，一个是初始化的，本质上不一样，如何区分？----用到前面的if的两段式
	*/
	t.Log(map1[1]) //0
	t.Log(map2[1]) //0
	if value, error := map2[1]; error {
		t.Logf("exist,value=%d", value)
	} else {
		t.Log("not exist")
	}
	if value, error := map1[1]; error {
		t.Logf(" exist,value=%d", value)
	} else {
		t.Log("not exist")
	}
}

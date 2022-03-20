package Arry_test

import "testing"

/**
数组初始化
*/
func TestArrInit(t *testing.T) {
	var arr1 [3]int                               //声明但是不初始化，系统会自动初始化为全0
	arr2 := [4]int{1, 2, 3, 4}                    //声明的同时进行初始化
	arr3 := [...]int{1, 5, 6, 7, 8}               //初始化的时候可以不写个数
	arr4 := [2][2]int{{1, 2}, {3, 4}}             //初始化一个二维数组
	arr5 := [2][4]int{{1, 2, 7, 9}, {5, 3, 4, 7}} //初始化一个二维数组
	t.Log(arr1, arr2, arr3, arr4, arr5)
}

/**
数组的遍历，可以用类似于foreach的结构
用range关键词给idx和elme赋值
*/
func TestArrTravel(t *testing.T) {
	arr2 := [4]int{1, 2, 3, 4}
	for i := 0; i < len(arr2); i++ {
		t.Log(arr2[i])
	}
	for idx, elem := range arr2 { //用range关键词给idx和elem赋值
		t.Log(idx, elem)
	}
	for _, elem := range arr2 { //_充当占位符
		t.Log(elem)
	}
}

/**
a[开始索引(包含) : 结束索引(不包含)]-----得到的是一个数组
头为空就从数组头开始，尾为空就是一直拷贝到最后一个等价于尾部填了len（arr）
*/
func TestArrSection(t *testing.T) {
	arr2 := [4]int{1, 2, 3, 4}
	arr3 := arr2[1:]          //从1开始一直到最后一个
	arr4 := arr2[1:len(arr2)] //和上面一句一样
	arr5 := arr2[:]           //全部拷贝
	arr6 := arr2[1:2]         //从1（包括）开始切片，不包括2
	arr7 := arr2[0:2]         //0-1 标全部拷贝过来
	arr8 := arr2[:2]          //从最开始一直拷贝到第二个，但是不包括第二个
	t.Log(arr2, arr3, arr4, arr5, arr6, arr7, arr8)
}

package slice_test

import "testing"

/**
1. slice是一种引用类型，它内部有一个指针指向数组。如果它利用了arr[x:y]来进行赋值，它的cap为len(arr)-x,它的len为y-x。由于是引用类型
所有它可以和其它的slice共用一个数组
2. slice的声明有两种。初始化有两种，一种是通过引用数组或者是数组的截取，另一种是通过make函数，还可以用append来进行初始化。
3. slice不能比较！！！！
*/
func TestSliceInit(t *testing.T) {
	a := [...]int{1, 2, 3, 4, 5, 6}
	var s1 []int  //声明1
	s2 := []int{} //声明2
	var s3 []int
	var s4 []int

	s1 = a[1:]  //初始化
	s2 = a[2:5] //初始化

	t.Log(s1, s2)
	t.Log(len(s1), cap(s1))
	t.Log(len(s2), cap(s2))

	s3 = append(s3, 1) //append的返回值是一个切片
	s4 = make([]int, 2, 3)
	t.Log(len(s3), cap(s3))
	t.Log(len(s4), cap(s4))
	s3 = append(s3, 1)
	s4 = append(s4, 3)
	t.Log(len(s3), cap(s3))
	t.Log(len(s4), cap(s4))
}

/**
可以看到通过append函数对slice进行增添元素，cap不够用的时候是会按照一定规律增加的，每次都是前面的2倍
*/
func TestSliceGrow_1(t *testing.T) {
	var s []int
	for i := 0; i < 10; i++ {
		s = append(s, i+1)
		t.Log(len(s), cap(s))
	}
}
func TestSliceGrow_2(t *testing.T) {
	var s []int = make([]int, 3, 5)
	for i := 0; i < 10; i++ {
		s = append(s, i+1)
		t.Log(len(s), cap(s))
	}
}

//可以看到两个slice引用了同一个数组，数组中的一个元素改变可能会影响到两个slice
func TestSliceSharingMemory(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var s1 []int = arr[2:8]
	var s2 []int = arr[4:9]
	t.Log(arr, s1, s2)
	arr[6] = 999
	t.Log(arr, s1, s2)
}

package arrayt

import "testing"

func TestArray(t *testing.T) {
	// 先声明, 后赋值
	var arr1 [3]string
	arr1[0] = "h"
	arr1[1] = "o"
	arr1[2] = "p"
	t.Logf("arr1 value %v", arr1)

	// 声明初始化同时操作
	arr2 := [3]int{1, 2, 3}
	t.Log(arr2, len(arr2), cap(arr2))

	// 不限定长度的声明
	arr3 := [...]int{1, 2, 4, 5, 6}
	t.Log(arr3, len(arr3), cap(arr3))

	// 遍历
	for idx, item := range arr3 {
		t.Logf("index: %v, value: %v", idx, item)
	}
}
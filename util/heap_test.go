package util

import "testing"

func TestMinHeap(t *testing.T) {
	hp := InitHeap(10, 2)
	hp.BuildHeapMin([]int64{2, 5, 1, 4, 3, 6, 10, 7, 9, 8}, 10)

	var i int64
	res := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i = 0; i < 10; i++ {
		if val := hp.DeleteMin(); val != res[i] {
			t.Errorf("response does not match:%v", val)
		}
	}
}

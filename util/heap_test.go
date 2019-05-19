package util

import (
	"testing"
)

func TestMinHeap(t *testing.T) {
	hp := InitHeap(11, 2)
	hp.BuildHeapMin([]int64{2, 5, 1, 4, 3, 6, 7, 9, 10}, 10)

	hp.InsertMin(8)
	if val := hp.DeleteMin(); val != 1 {
		t.Errorf("response does not match:%v", val)
	}

	var i int64
	res := []int64{2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i = 0; i < 9; i++ {
		if val := hp.DeleteMin(); val != res[i] {
			t.Errorf("response does not match:%v", val)
		}
	}

	hp.DeleteMin()
	hp.DeleteMin()

	if val := hp.DeleteMin(); val != -1 {
		t.Errorf("response does not match:%v", val)
	}
}

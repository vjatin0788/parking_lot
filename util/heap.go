package util

import "math"

type HeapNode struct {
	Arr      []int64
	Capacity int64 //size of heap.
	Count    int64 //number of elements.
	HeapType int64
}

func (node *HeapNode) Parent(i int64) int64 {

	pos := (i - 1) / 2
	if (i-1) >= 0 && pos < int64(len(node.Arr)) {
		return pos
	}
	return -1
}

func (node *HeapNode) Child(i int64) (int64, int64) {
	firstChild := 2*i + 1
	secondChild := 2*i + 2

	if firstChild < node.Count && secondChild < node.Count {
		return firstChild, secondChild
	}
	if firstChild < node.Count && secondChild >= node.Count {
		return firstChild, -1
	}
	if firstChild >= node.Count && secondChild < node.Count {
		return -1, secondChild
	}

	return -1, -1

}

func (node *HeapNode) GetMaxOrMin() int64 {
	if node.Count > 0 {
		return node.Arr[0]
	}
	return -1
}

func (node *HeapNode) PercolateDownMin(i int64) {

	if i < 0 {
		return
	}

	var (
		l, r, min int64
	)

	l, r = node.Child(i)
	if l != -1 && node.Arr[l] < node.Arr[i] {
		min = l
	} else {
		min = i
	}
	if r != -1 && node.Arr[r] < node.Arr[min] {
		min = r
	}

	if min != i {
		temp := node.Arr[min]
		node.Arr[min] = node.Arr[i]
		node.Arr[i] = temp
	}

	if min == i {
		return
	}

	node.PercolateDownMin(min)

}

func (node *HeapNode) DeleteMin() int64 {
	if node == nil {
		return -1
	}
	if node.Count == 0 {
		return -1
	}
	removedElement := node.Arr[0]
	node.Arr[0] = node.Arr[node.Count-1]
	node.Arr[node.Count-1] = math.MaxInt64
	node.Count -= 1
	node.PercolateDownMin(0)
	return removedElement
}

func (node *HeapNode) InsertMin(data int64) {
	if node.Count == node.Capacity {
		node.ReSize(2 * node.Capacity)
	}
	node.Arr[node.Count] = data
	node.Count++
	node.PercolateUpMin(node.Count - 1)
}

func (node *HeapNode) ReSize(num int64) {
	var newArr []int64
	newArr = make([]int64, num)

	for idx := 0; idx < int(num); idx++ {
		if idx < len(node.Arr) {
			newArr[idx] = node.Arr[idx]
		} else {
			newArr[idx] = math.MaxInt64
		}
	}
	node.Capacity = num
	node.Arr = newArr
}

func (node *HeapNode) PercolateUpMin(i int64) {
	var (
		p int64
	)
	p = node.Parent(i)
	for p >= 0 {
		if node.Arr[p] > node.Arr[i] {
			temp := node.Arr[p]
			node.Arr[p] = node.Arr[i]
			node.Arr[i] = temp
		} else {
			break
		}
		i = p
		p = node.Parent(p)

	}
}

func (node *HeapNode) BuildHeapMin(arr []int64, num int64) {
	if node == nil {
		return
	}
	var idx int64
	if node.Capacity < num {
		node.ReSize(num)
	}
	for idx := range arr {
		node.Arr[idx] = arr[idx]
	}
	node.Count = num

	for idx = (num - 1) / 2; idx >= 0; idx-- {
		node.PercolateDownMin(idx - 1)
	}
}

func InitHeap(capacity int64, hType int64) *HeapNode {
	arr := make([]int64, capacity)

	for idx := range arr {
		arr[idx] = math.MaxInt64
	}

	return &HeapNode{
		Capacity: capacity,
		Arr:      arr,
		HeapType: hType, // 1 := max heap type
	}
}

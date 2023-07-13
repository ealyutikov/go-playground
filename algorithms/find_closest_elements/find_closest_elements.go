package find_closest_elements

import (
	"container/heap"
	"math"
)

func findClosestElements(arr []int, k int, x int) []int {
	if len(arr) < k {
		return arr
	}

	minHeap := intHeap{}
	heap.Init(&minHeap)

	for i := 0; i < len(arr); i++ {
		if len(minHeap) < k {
			heap.Push(&minHeap, arr[i])
		} else {
			min := heap.Pop(&minHeap).(int)
			if math.Abs(float64(min-x)) < math.Abs(float64(arr[i]-x)) ||
				(math.Abs(float64(min-x)) == math.Abs(float64(arr[i]-x)) && min < arr[i]) {
				heap.Push(&minHeap, min)
			} else {
				heap.Push(&minHeap, arr[i])
			}
		}
	}

	result := make([]int, 0, k)
	for minHeap.Len() > 0 {
		result = append(result, heap.Pop(&minHeap).(int))
	}

	return result
}

// https://www.sohamkamani.com/golang/heap/
type intHeap []int

func (m intHeap) Len() int {
	return len(m)
}

func (m intHeap) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m intHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (h *intHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

package sol

import "container/heap"

type IntMinHeap []int

// Len() int
func (h *IntMinHeap) Len() int {
	return len(*h)
}

// Swap(i, j int)
func (h *IntMinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

// Less(i, j int) bool
func (h *IntMinHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

// Push(val interface{})
func (h *IntMinHeap) Push(val interface{}) {
	*h = append(*h, val.(int))
}

// Pop() interface{}
func (h *IntMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type KthLargest struct {
	k             int
	priorityQueue *IntMinHeap
}

func Constructor(k int, nums []int) KthLargest {
	priorityQueue := &IntMinHeap{}
	heap.Init(priorityQueue)
	for idx := 0; idx < len(nums); idx++ {
		heap.Push(priorityQueue, nums[idx])
	}
	return KthLargest{
		k:             k,
		priorityQueue: priorityQueue,
	}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.priorityQueue, val)
	for this.priorityQueue.Len() > this.k {
		heap.Pop(this.priorityQueue)
	}
	return (*this.priorityQueue)[0]
}

/**
* Your KthLargest object will be instantiated and called as such:
* obj := Constructor(k, nums);
* param_1 := obj.Add(val);
 */

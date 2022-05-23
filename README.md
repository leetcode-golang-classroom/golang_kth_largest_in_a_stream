# golang_kth_largest_in_a_stream

Design a class to find the $`k^{th}`$ largest element in a stream. Note that it is the $`k^{th}`$ largest element in the sorted order, not the $`k^{th}`$ distinct element.

Implement `KthLargest` class:

- `KthLargest(int k, int[] nums)` Initializes the object with the integer `k` and the stream of integers `nums`.
- `int add(int val)` Appends the integer `val` to the stream and returns the element representing the $`k^{th}`$ largest element in the stream.

## Examples

**Example 1:**

```
Input
["KthLargest", "add", "add", "add", "add", "add"]
[[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
Output
[null, 4, 5, 5, 8, 8]

Explanation
KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
kthLargest.add(3);   // return 4
kthLargest.add(5);   // return 5
kthLargest.add(10);  // return 5
kthLargest.add(9);   // return 8
kthLargest.add(4);   // return 8

```

**Constraints:**

- `1 <= k <= $10^4$`
- `0 <= nums.length <= $10^4$`
- $`-10^4$ <= nums[i] <= $10^4$`
- $`-10^4$ <= val <= $10^4$`
- At most $`10^4`$ calls will be made to `add`.
- It is guaranteed that there will be at least `k` elements in the array when you search for the $`k^{th}`$ element.

## 解析

給定一個正整數 k 還有整數陣列 nums 

求實作一個資料結構 `KthLargest` 具有以下 method

1. Constructor(k int, nums []int): 回傳 `KthLargest`  其中 k 代表每次回傳第 k 大的數值, nums 代表傳入的所有資料
2. add(val int) int : 在nums 加入 val 後，回傳當下第 k 的數值

直覺的作法是用一個 int 儲存 k , 還有一個陣列儲存 nums 

對 nums 做 sort 如果是 heap sort 則需要 O(nlogn)

然後每次都把加入的值對前 k 大的值做比較 O(k)

如果要優化比較的值

可以考慮用 min-heap 儲存第一次輸入也是 nlogn

然後每次只保留 k 個 也就是 前 n-k + 1 個會小會被移除

所有留下來在 heap 最前面的值就是第 n-k 小 也就是第 k 大

每次輸入新的值都是 log(k)

 

所以比第1種方法用 min-heap 是一個比較優化的方式

![](https://i.imgur.com/tnogT2C.png)

## 程式碼

```go
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

```
## 困難點

1. 理解 Priority Queue 在這題做到的排序功能

## Solve Point

- [x]  Understand what problem need to solve
- [x]  Analysis Complexity
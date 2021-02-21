package sorting

// maxHeap
type simpleHeap struct {
	values []int
}

func (h *simpleHeap) getLeftChild(i int) int {
	leftInd := 2*i + 1
	if len(h.values) <= leftInd {
		return -1
	}
	return h.values[leftInd]
}

func (h *simpleHeap) getLeftChildInd(i int) int {
	return 2*i + 1
}

func (h *simpleHeap) getRightChild(i int) int {
	rightInd := 2*i + 2
	if len(h.values) <= rightInd {
		return -1
	}
	return h.values[rightInd]
}

func (h *simpleHeap) getRightChildInd(i int) int {
	return 2*i + 2
}

func (h *simpleHeap) getParentInd(i int) int {
	parentInd := (i - 1) / 2
	if parentInd < 0 {
		return -1
	}
	return parentInd
}

func (h *simpleHeap) swap(i, j int) {
	h.values[i], h.values[j] = h.values[j], h.values[i]
}

func (h *simpleHeap) push(val int) {
	h.values = append(h.values, val)
	h.heapifyUp(len(h.values) - 1)
}

func (h *simpleHeap) pop() int {
	val := h.values[0]
	h.values[0] = h.values[len(h.values)-1]
	h.values = h.values[:len(h.values)-1]
	h.heapifyDown(0)
	return val
}

// O(logn)
func (h *simpleHeap) heapifyUp(i int) {
	parentInd := h.getParentInd(i)
	for parentInd != -1 && h.values[parentInd] < h.values[i] {
		h.swap(parentInd, i)
		i = parentInd
		parentInd = h.getParentInd(i)
	}
}

// O(logn)
func (h *simpleHeap) heapifyDown(i int) {
	swapChildInd := h.getLeftChildInd(i)
	for h.getLeftChild(i) != -1 && h.values[swapChildInd] > h.values[i] {
		// chose the greatest child
		if h.getRightChild(i) != -1 && h.getRightChild(i) > h.getLeftChild(i) {
			swapChildInd = h.getRightChildInd(i)
		}
		h.swap(swapChildInd, i)
		i = swapChildInd
		swapChildInd = h.getLeftChild(i)
	}
}

func HeapSort(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	maxHeap := simpleHeap{values: make([]int, 0, len(nums))}
	// O(nlogn)
	for _, n := range nums {
		maxHeap.push(n)
	}

	sorted := make([]int, 0, len(nums))
	// O(nlogn)
	for i := 0; i < len(nums); i++ {
		sorted = append(sorted, maxHeap.pop())
	}

	return sorted
}

func KGreat(nums []int, k int) int {
	if k > len(nums) {
		return -1
	}

	maxHeap := simpleHeap{values: make([]int, 0, len(nums))}
	// O(nlogn)
	for _, n := range nums {
		maxHeap.push(n)
	}

	// O(klogn)
	for i := 0; i < k-1; i++ {
		maxHeap.pop()
	}
	return maxHeap.pop()
}

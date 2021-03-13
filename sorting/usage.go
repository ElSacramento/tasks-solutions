package sorting

import (
	"container/heap"
	"sort"
)

type stack []int

func (s *stack) pop() int {
	old := *s
	n := len(old)
	elem := old[n-1]
	old[n-1] = -1
	*s = old[:n-1]
	return elem
}

func (s *stack) peak() int {
	st := *s
	return st[len(st)-1]
}

func (s *stack) add(a int) {
	old := *s
	old = append(old, a)
	*s = old
}

type sortedInfo struct {
	sort.IntSlice
	idx []int
}

func (s sortedInfo) Swap(i, j int) {
	s.IntSlice.Swap(i, j)
	s.idx[i], s.idx[j] = s.idx[j], s.idx[i]
}

// leetcode: 975
// monotonic stack approach + DP from the end of the array
func oddEvenJumps(arr []int) int {
	n := len(arr)

	sortedASC := sortedInfo{
		IntSlice: make(sort.IntSlice, n),
		idx:      make([]int, n),
	}
	sortedDESC := sortedInfo{
		IntSlice: make(sort.IntSlice, n),
		idx:      make([]int, n),
	}
	for i := 0; i < n; i++ {
		sortedASC.idx[i] = i
		sortedDESC.idx[i] = i

		sortedASC.IntSlice[i] = arr[i]
		sortedDESC.IntSlice[i] = -arr[i] // to sort in desc order
	}
	sort.Stable(sortedASC)
	sort.Stable(sortedDESC)

	oddNext := make([]int, n)
	evenNext := make([]int, n)
	for i := 0; i < n; i++ {
		oddNext[i] = -1
		evenNext[i] = -1
	}

	simpleStack := make(stack, 0)

	for _, ind := range sortedASC.idx {
		for len(simpleStack) != 0 && ind > simpleStack.peak() {
			oddNext[simpleStack.pop()] = ind
		}
		simpleStack.add(ind)
	}

	for len(simpleStack) != 0 {
		simpleStack.pop()
	}

	for _, ind := range sortedDESC.idx {
		for len(simpleStack) != 0 && ind > simpleStack.peak() {
			evenNext[simpleStack.pop()] = ind
		}
		simpleStack.add(ind)
	}

	// DP
	odd := make([]bool, n)
	even := make([]bool, n)
	odd[n-1], even[n-1] = true, true

	for i := len(arr) - 2; i >= 0; i-- {
		if oddNext[i] != -1 {
			odd[i] = even[oddNext[i]]
		}
		if evenNext[i] != -1 {
			even[i] = odd[evenNext[i]]
		}
	}

	counter := 0
	for _, v := range odd {
		if v {
			counter++
		}
	}
	return counter
}

type Item struct {
	upValue   int32
	downValue int32
}

type Heap []Item

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	change1 := float64(h[i].upValue+1)/float64(h[i].downValue+1) - float64(h[i].upValue)/float64(h[i].downValue)
	change2 := float64(h[j].upValue+1)/float64(h[j].downValue+1) - float64(h[j].upValue)/float64(h[j].downValue)
	return change1 > change2
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x interface{}) {
	item := x.(Item)
	*h = append(*h, item)
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

// calculate the minimum number of additional five-star reviews the company needs to meet the threshold
// amazon OA
// O(nlogn) + space for heap
func fiveStarReviews(productRatings [][]int32, ratingThreshold int32) int32 {
	n := len(productRatings)
	var sum float64

	values := make([]Item, 0)
	for i := 0; i < n; i++ {
		// if 5 star reviews number is equal to total, adding one more 5 star review won't change anything
		if productRatings[i][0] == productRatings[i][1] {
			sum += 1
			continue
		}
		values = append(values, Item{upValue: productRatings[i][0], downValue: productRatings[i][1]})

		sum += float64(productRatings[i][0]) / float64(productRatings[i][1])
	}
	ratingsHeap := Heap(values)
	heap.Init(&ratingsHeap) // O(nlogn) + space O(n)

	threshHold := float64(ratingThreshold) / 100 * float64(n)
	changes := 0
	for sum < threshHold {
		item := heap.Pop(&ratingsHeap).(Item)
		sum += float64(item.upValue+1)/float64(item.downValue+1) - float64(item.upValue)/float64(item.downValue)
		changes++

		heap.Push(&ratingsHeap, Item{upValue: item.upValue + 1, downValue: item.downValue + 1})
	}
	return int32(changes)
}

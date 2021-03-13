package sorting

import (
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
		IntSlice: make(sort.IntSlice, n, n),
		idx:      make([]int, n, n),
	}
	sortedDESC := sortedInfo{
		IntSlice: make(sort.IntSlice, n, n),
		idx:      make([]int, n, n),
	}
	for i := 0; i < n; i++ {
		sortedASC.idx[i] = i
		sortedDESC.idx[i] = i

		sortedASC.IntSlice[i] = arr[i]
		sortedDESC.IntSlice[i] = -arr[i] // to sort in desc order
	}
	sort.Stable(sortedASC)
	sort.Stable(sortedDESC)

	oddNext := make([]int, n, n)
	evenNext := make([]int, n, n)
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
	odd := make([]bool, n, n)
	even := make([]bool, n, n)
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

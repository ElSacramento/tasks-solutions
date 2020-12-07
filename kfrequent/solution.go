package kfrequent

import (
	"container/heap"
)

// must be less than O(nlogn), cant use heapSort
func topKFrequentNums(nums []int, k int) []int {
	counters := make(map[int]int, 0)
	maxCount := 0
	for _, v := range nums {
		if _, found := counters[v]; found {
			counters[v]++
			if counters[v] > maxCount {
				maxCount = counters[v]
			}
			continue
		}
		counters[v] = 1
		if counters[v] > maxCount {
			maxCount = counters[v]
		}
	}

	buckets := make([][]int, maxCount+1, maxCount+1)
	for n, count := range counters {
		items := buckets[count]
		if len(items) == 0 {
			buckets[count] = []int{n}
			continue
		}
		items = append(items, n)
		buckets[count] = items
	}

	result := make([]int, 0, k)
	for i := len(buckets) - 1; i >= 0; i-- {
		if len(buckets[i]) != 0 {
			result = append(result, buckets[i]...)
			if len(result) == k {
				return result
			}
		}
	}
	return result
}

type wordInfo struct {
	value string
	count int
}

type wordsHeap struct {
	words []*wordInfo
}

func (w wordsHeap) Len() int {
	return len(w.words)
}

func (w wordsHeap) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest
	if w.words[i].count > w.words[j].count {
		return true
	}
	if w.words[i].count == w.words[j].count {
		return w.words[i].value < w.words[j].value
	}
	return false
}

func (w wordsHeap) Swap(i, j int) {
	w.words[i], w.words[j] = w.words[j], w.words[i]
}

func (w *wordsHeap) Push(x interface{}) {
	elem := x.(*wordInfo)
	w.words = append(w.words, elem)
}

func (w *wordsHeap) Pop() interface{} {
	n := len(w.words)
	elem := w.words[n-1]
	w.words = w.words[0 : n-1]
	return elem
}

// heapSort - O(nlogk)
func topKFrequentWords(words []string, k int) []string {
	counts := make(map[string]int, 0)

	for _, s := range words {
		counts[s] += 1
	}

	wHeap := wordsHeap{words: make([]*wordInfo, 0)}
	heap.Init(&wHeap)
	for w, c := range counts {
		heap.Push(&wHeap, &wordInfo{value: w, count: c})
	}

	response := make([]string, 0, k)
	for i := 0; i < k; i++ {
		elem := heap.Pop(&wHeap)
		word := elem.(*wordInfo)
		response = append(response, word.value)
	}
	return response
}

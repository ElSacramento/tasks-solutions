package kfrequent

import (
	"container/heap"
	"strings"
)

// must be less than O(nlogn), cant use heapSort
// leetcode: 347
func topKFrequentNums(nums []int, k int) []int {
	counters := make(map[int]int)
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

	buckets := make([][]int, maxCount+1)
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
// leetcode: 692
func topKFrequentWords(words []string, k int) []string {
	counts := make(map[string]int)

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

// find most frequent word, that isn't banned - O(n)
// space usage can be reduced by removing strings.FieldsFunc
// leetcode: 819
func mostCommonWord(paragraph string, banned []string) string {
	words := strings.FieldsFunc(paragraph, func(r rune) bool {
		return r == '!' || r == '?' || r == '\'' || r == ',' || r == ';' || r == '.' || r == ' '
	})
	for i := 0; i < len(words); i++ {
		words[i] = strings.ToLower(words[i])
	}

	freq := make(map[string]int)
	for _, s := range words {
		if _, found := freq[s]; found {
			freq[s]++
			continue
		}
		freq[s] = 1
	}

	for _, s := range banned {
		freq[s] = 0
	}

	maxCount := 0
	maxWord := ""
	for word, count := range freq {
		if count > maxCount {
			maxCount = count
			maxWord = word
		}
	}

	return maxWord
}

type symbolInfo struct {
	count int
	value rune
}

type heapInfo struct {
	symbols []*symbolInfo
}

func (h heapInfo) Len() int {
	return len(h.symbols)
}

func (h heapInfo) Swap(i, j int) {
	h.symbols[i], h.symbols[j] = h.symbols[j], h.symbols[i]
}

func (h heapInfo) Less(i, j int) bool {
	return h.symbols[i].count > h.symbols[j].count
}

func (h *heapInfo) Push(x interface{}) {
	elem := x.(*symbolInfo)
	h.symbols = append(h.symbols, elem)
}

func (h *heapInfo) Pop() interface{} {
	n := len(h.symbols)
	elem := h.symbols[n-1]
	h.symbols = h.symbols[:n-1]
	return elem
}

// O(nlogn) , space a lot
// same symbols can't be close
func reorganizeString(S string) string {
	if len(S) == 1 {
		return S
	}

	// space O(n), complexity O(n)
	frequency := make(map[rune]int)
	for _, s := range S {
		frequency[s]++
	}

	// space O(n), complexity O(n)
	symbols := make([]*symbolInfo, 0, len(frequency))
	for symbolRune, symbolCount := range frequency {
		symbols = append(symbols, &symbolInfo{count: symbolCount, value: symbolRune})
	}

	sHeap := heapInfo{symbols: symbols}
	// complexity O(n)
	heap.Init(&sHeap)

	// space O(n)
	var prev rune
	result := make([]rune, 0)
	// complexity O(n * 5logn)
	for sHeap.Len() != 0 {
		// complexity O(logn)
		elem := heap.Pop(&sHeap).(*symbolInfo)
		if prev == 0 || (prev != 0 && prev != elem.value) {
			result = append(result, elem.value)
			prev = elem.value
			elem.count--
			// complexity O(logn)
			if elem.count != 0 {
				heap.Push(&sHeap, elem)
			}
			continue
		}
		// prev == elem.Value
		if sHeap.Len() == 0 {
			// impossible to find another symbol
			return ""
		}

		// complexity O(logn)
		nextElem := heap.Pop(&sHeap).(*symbolInfo)
		result = append(result, nextElem.value)
		prev = nextElem.value
		nextElem.count--
		// complexity O(logn)
		if nextElem.count != 0 {
			heap.Push(&sHeap, nextElem)
		}
		// complexity O(logn)
		heap.Push(&sHeap, elem)
	}

	// complexity O(n)
	var response string
	for _, el := range result {
		response += string(el)
	}

	return response
}

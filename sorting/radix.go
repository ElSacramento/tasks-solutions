package sorting

import (
	"fmt"
)

// O(n + base)
func CountingSort(nums []int) []int {
	base := 3 // because we can only have 0, 1, 2
	frequency := make([]int, base)

	// O(n)
	for _, el := range nums {
		frequency[el]++
	}
	fmt.Printf("%p\n", frequency)

	indexes := frequency // reuse slice, no additional space
	// O(base)
	for i := 0; i < len(indexes); i++ {
		if i == 0 {
			continue
		}
		indexes[i] = indexes[i] + indexes[i-1]
	}
	// O(base)
	for i := len(indexes) - 1; i > 0; i-- {
		indexes[i] = indexes[i-1]
	}
	indexes[0] = 0
	fmt.Printf("%p\n", indexes)

	result := make([]int, len(nums))
	// O(n)
	for _, el := range nums {
		ind := indexes[el]
		result[ind] = el
		indexes[el]++
	}
	return result
}

func sortColors(nums []int) {
	base := 3                      // because we can only have 0, 1, 2
	frequency := make([]int, base) // constant space

	// O(n)
	for _, el := range nums {
		frequency[el]++
	}

	start := 0
	// O(base * n)
	for i := 0; i < len(frequency); i++ {
		count := frequency[i] + start
		for j := start; j < count; j++ {
			nums[j] = i
		}
		start += frequency[i]
	}
}

// space: O(length + base + n + C) ----> O(n)
// complexity: O(length + length * (3n + 3base)) = O(length * 3 (n + base)) ------> O(length * (n + base))
func RadixSort(nums []int) {
	base := 10  // because 0, 1, ... 9
	length := 3 // because 0 <= elem < 1000

	// space O(length)
	powers := make([]int, length+1)
	powers[0] = 1
	// O(length)
	for i := 1; i <= length; i++ {
		powers[i] = powers[i-1] * base
	}

	// space O(base)
	frequency := make([]int, base)
	var indexes []int

	// space O(n)
	tempo := make([]int, len(nums))

	// O(length)
	for i := 1; i <= length; i++ {
		// O(n)
		for _, el := range nums {
			part := el % powers[i] / powers[i-1]
			frequency[part]++
		}

		indexes = frequency
		// O(base)
		for j := 0; j < len(indexes); j++ {
			if j == 0 {
				continue
			}
			indexes[j] = indexes[j] + indexes[j-1]
		}

		// O(base)
		for j := len(indexes) - 1; j > 0; j-- {
			indexes[j] = indexes[j-1]
		}
		indexes[0] = 0

		// O(n)
		for _, el := range nums {
			part := el % powers[i] / powers[i-1]
			ind := indexes[part]
			tempo[ind] = el
			indexes[part]++
		}

		// O(n)
		copy(nums, tempo)

		// O(base)
		for i := range frequency {
			frequency[i] = 0
		}
	}
}

package subsequence

import (
	"sort"
)

// [10, 4, 5, 1, 2, 3, 7]
// [10]
// found greater element, replace, because the length is the same [4]
// [4, 5]
// [1, 5]
// [1, 2]
// [1, 2, 3]
// [1, 2, 3, 7]

// [10, 4, 5, 1, 7]
// [4, 5]
// [1, 5]
// [1, 5, 7]

func LengthOfLIS(nums []int) int {
	sorted := make([]int, 0)
	sorted = append(sorted, nums[0])

	// O(nlogn)
	for _, num := range nums {
		foundInd := sort.Search(len(sorted), func(k int) bool {
			return sorted[k] >= num
		})
		if foundInd != len(sorted) {
			sorted[foundInd] = num
		} else {
			sorted = append(sorted, num)
		}
	}
	return len(sorted)
}

package other

import (
	"math"
)

// leetcode: 126
func findPeakElement(nums []int) int {
	infinity := math.MaxInt32 + 10
	left, right := -infinity, -infinity
	i, j := 0, len(nums)-1
	// binary search = O(logn)
	for i < j {
		k := (i + j) / 2
		if k == 0 {
			left = -infinity
		} else {
			left = nums[k-1]
		}
		if k >= len(nums)-1 {
			right = -infinity
		} else {
			right = nums[k+1]
		}
		current := nums[k]

		// it's a peak!
		if current > left && current > right {
			return k
		}
		// going up
		if current > left && current < right {
			i = k + 1
			continue
		}
		// going down
		if current < left && current > right {
			j = k
			continue
		}
		// it's a bottom, let's go right
		if current < left && current < right {
			i = k + 1
			continue
		}
	}
	return i
}

package sorting

func qsort(nums []int, left, right int) {
	if left >= right {
		return
	}
	index := partition(nums, left, right)
	qsort(nums, left, index)
	qsort(nums, index+1, right)
}

func partition(nums []int, left, right int) int {
	pivot := nums[left+(right-left)/2]
	for {
		if nums[left] < pivot {
			left++
			continue
		}
		if nums[right] > pivot {
			right--
			continue
		}
		if left >= right {
			return right
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
}

// O(nlogn)
// worst case O(n*n), if pivot is first or last element
func quickSort(nums []int) {
	qsort(nums, 0, len(nums)-1)
}

// O(n*n)
func insertionSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		current := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > current {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = current
	}
}

// O(nlogn) + space O(n)
func mergeSort(nums []int) []int {
	return []int{}
}

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

func mergeSortRec(nums []int, start, end int) []int {
	// parts sort
	length := end - start + 1
	if length == 1 {
		return []int{nums[start]}
	}
	if length == 2 {
		if nums[start] <= nums[end] {
			return []int{nums[start], nums[end]}
		}
		return []int{nums[end], nums[start]}
	}

	middle := length / 2
	firstPart := mergeSortRec(nums, start, start+middle-1)
	secondPart := mergeSortRec(nums, start+middle, end)

	// parts merge
	result := make([]int, 0, len(firstPart)+len(secondPart))
	fInd, sInd := 0, 0
	for fInd != len(firstPart) && sInd != len(secondPart) {
		if firstPart[fInd] <= secondPart[sInd] {
			result = append(result, firstPart[fInd])
			fInd++
			continue
		}
		result = append(result, secondPart[sInd])
		sInd++
	}
	if fInd != len(firstPart) {
		result = append(result, firstPart[fInd:]...)
	} else if sInd != len(secondPart) {
		result = append(result, secondPart[sInd:]...)
	}
	if len(result) != cap(result) {
		panic("impossible")
	}
	return result
}

// O(nlogn) + space O(n)
func mergeSort(nums []int) []int {
	return mergeSortRec(nums, 0, len(nums)-1)
}

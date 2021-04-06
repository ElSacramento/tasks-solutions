package other

type stack []int

func (s *stack) add(v int) {
	old := *s
	old = append(old, v)
	*s = old
}

func (s *stack) pop() int {
	old := *s
	n := len(old)
	value := old[n-1]
	*s = old[:n-1]
	return value
}

func (s *stack) peak() int {
	old := *s
	n := len(old)
	return old[n-1]
}

// linear time O(n)
// leetcode: 84
func largestRectangleArea(heights []int) int {
	indexesToCalculate := make(stack, 0)

	resultMax := 0
	i := 0
	var width int
	// O(n)
	for i < len(heights) {
		// monotonic stack , height[x1] <= height[x2]
		if len(indexesToCalculate) == 0 || heights[i] >= heights[indexesToCalculate.peak()] {
			indexesToCalculate.add(i)
			i++
			continue
		}

		localMaxHeightIndex := indexesToCalculate.pop()

		if len(indexesToCalculate) == 0 {
			width = i
		} else {
			// before the prev element
			nextIndex := indexesToCalculate.peak()
			width = i - (nextIndex + 1)
		}

		area := heights[localMaxHeightIndex] * width
		if area > resultMax {
			resultMax = area
		}
	}

	if len(indexesToCalculate) == 0 {
		return resultMax
	}

	// O(n)
	for len(indexesToCalculate) != 0 {
		localMaxHeightIndex := indexesToCalculate.pop()

		if len(indexesToCalculate) == 0 {
			width = i
		} else {
			// before the prev element
			nextIndex := indexesToCalculate.peak()
			width = i - (nextIndex + 1)
		}

		area := heights[localMaxHeightIndex] * width
		if area > resultMax {
			resultMax = area
		}
	}
	return resultMax
}

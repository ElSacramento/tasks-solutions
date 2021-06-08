package sorting

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBasicSorts(t *testing.T) {
	type testCase struct {
		nums     []int
		expected []int
	}
	cases := []testCase{
		{
			nums:     []int{5, 1, 3, 7, 8, 4},
			expected: []int{1, 3, 4, 5, 7, 8},
		},
		{
			nums:     []int{3, 7, 4, 1},
			expected: []int{1, 3, 4, 7},
		},
		{
			nums:     []int{5, 2, 3, 1},
			expected: []int{1, 2, 3, 5},
		},
		{
			nums:     []int{1, 2, 3, 4},
			expected: []int{1, 2, 3, 4},
		},
		{
			nums:     []int{4, 3, 2, 1},
			expected: []int{1, 2, 3, 4},
		},
	}
	for _, elem := range cases {
		t.Run("qsort", func(t *testing.T) {
			quickSort(elem.nums)
			require.Equal(t, elem.expected, elem.nums)
		})
		t.Run("insertion sort", func(t *testing.T) {
			insertionSort(elem.nums)
			require.Equal(t, elem.expected, elem.nums)
		})
	}
}

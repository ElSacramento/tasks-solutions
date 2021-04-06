package other

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLargestRectangleArea(t *testing.T) {
	{
		heights := []int{2, 1, 4, 5, 3, 1}
		require.Equal(t, 9, largestRectangleArea(heights))
	}
	{
		heights := []int{2, 1, 4, 1, 3, 4}
		require.Equal(t, 6, largestRectangleArea(heights))
	}
	{
		heights := []int{2, 1, 4, 1, 3, 4, 4}
		require.Equal(t, 9, largestRectangleArea(heights))
	}
	{
		heights := []int{2, 1, 0, 3}
		require.Equal(t, 3, largestRectangleArea(heights))
	}
	{
		heights := []int{2, 1, 5, 6, 2, 3}
		require.Equal(t, 10, largestRectangleArea(heights))
	}
	{
		heights := []int{2, 4}
		require.Equal(t, 4, largestRectangleArea(heights))
	}
	{
		heights := []int{2, 1, 4, 1, 3}
		require.Equal(t, 5, largestRectangleArea(heights))
	}
}

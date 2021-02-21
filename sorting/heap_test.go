package sorting

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHeapSort(t *testing.T) {
	{
		nums := []int{4, 1, 5, 8, 2, 3, 9}
		require.Equal(t, []int{9, 8, 5, 4, 3, 2, 1}, HeapSort(nums))
	}
	{
		nums := []int{1, 2, 3}
		require.Equal(t, []int{3, 2, 1}, HeapSort(nums))
	}
	{
		nums := []int{4}
		require.Equal(t, []int{4}, HeapSort(nums))
	}
}

func TestKGreat(t *testing.T) {
	{
		nums := []int{4, 1, 5, 8, 2, 3, 9}
		require.Equal(t, 5, KGreat(nums, 3))
	}
	{
		nums := []int{4, 1, 5, 8, 2, 3, 9}
		require.Equal(t, 2, KGreat(nums, 6))
	}
	{
		nums := []int{1, 2, 3}
		require.Equal(t, 2, KGreat(nums, 2))
	}
	{
		nums := []int{4}
		require.Equal(t, 4, KGreat(nums, 1))
	}
	{
		nums := []int{4}
		require.Equal(t, -1, KGreat(nums, 3))
	}
}

package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountingSort(t *testing.T) {
	{
		arr := []int{2, 0, 0, 1, 2, 1}
		assert.Equal(t, []int{0, 0, 1, 1, 2, 2}, CountingSort(arr))
	}
	{
		arr := []int{2, 0}
		assert.Equal(t, []int{0, 2}, CountingSort(arr))
	}
	{
		arr := []int{0, 1, 2}
		assert.Equal(t, []int{0, 1, 2}, CountingSort(arr))
	}
	{
		arr := []int{2}
		assert.Equal(t, []int{2}, CountingSort(arr))
	}
}

func TestSortColors(t *testing.T) {
	{
		arr := []int{2, 0, 0, 1, 2, 1}
		sortColors(arr)
		assert.Equal(t, []int{0, 0, 1, 1, 2, 2}, arr)
	}
	{
		arr := []int{2, 0}
		sortColors(arr)
		assert.Equal(t, []int{0, 2}, arr)
	}
	{
		arr := []int{0, 1, 2}
		sortColors(arr)
		assert.Equal(t, []int{0, 1, 2}, arr)
	}
	{
		arr := []int{2}
		assert.Equal(t, []int{2}, CountingSort(arr))
	}
}

func TestRadixSort(t *testing.T) {
	{
		arr := []int{201, 10, 34, 138, 99}
		RadixSort(arr)
		assert.Equal(t, []int{10, 34, 99, 138, 201}, arr)
	}
	{
		arr := []int{201, 200, 170, 17, 3}
		RadixSort(arr)
		assert.Equal(t, []int{3, 17, 170, 200, 201}, arr)
	}
	{
		arr := []int{209}
		RadixSort(arr)
		assert.Equal(t, []int{209}, arr)
	}
	{
		arr := []int{55, 89, 890}
		RadixSort(arr)
		assert.Equal(t, []int{55, 89, 890}, arr)
	}
}

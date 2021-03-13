package sorting

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOddEvenJumps(t *testing.T) {
	{
		arr := []int{10, 13, 12, 14, 15}
		require.Equal(t, 2, oddEvenJumps(arr))
	}
	{
		arr := []int{2, 3, 1, 1, 4}
		require.Equal(t, 3, oddEvenJumps(arr))
	}
	{
		arr := []int{5, 1, 3, 4, 2}
		require.Equal(t, 3, oddEvenJumps(arr))
	}
	{
		arr := []int{1, 2, 1, 2, 1}
		require.Equal(t, 4, oddEvenJumps(arr))
	}
	{
		arr := []int{1, 2, 1, 1, 1, 1}
		require.Equal(t, 5, oddEvenJumps(arr))
	}
	{
		arr := []int{1, 2, 3, 2, 1, 4, 4, 5}
		require.Equal(t, 6, oddEvenJumps(arr))
	}
	{
		arr := []int{8, 75, 87, 92, 84, 89, 92, 54, 39, 70}
		require.Equal(t, 8, oddEvenJumps(arr))
	}
}

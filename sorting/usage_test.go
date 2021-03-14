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

func TestFiveStarReviews(t *testing.T) {
	{
		ratings := [][]int32{
			{4, 4},
			{1, 2},
			{3, 6},
		}
		// the percentage for this seller is ((4 / 4) + (1/2) + (3/6))/3 = 66.66%
		require.Equal(t, int32(3), fiveStarReviews(ratings, 77))
	}
	{
		ratings := [][]int32{
			{1, 2},
			{1, 2},
		}
		require.Equal(t, int32(4), fiveStarReviews(ratings, 75))
	}
	{
		ratings := [][]int32{
			{2, 2},
			{1, 2},
		}
		require.Equal(t, int32(1), fiveStarReviews(ratings, 80))
	}
}

func TestKClosest(t *testing.T) {
	{
		points := [][]int{
			{3, 3},
			{5, -1},
			{-2, 7},
		}
		require.Equal(t, [][]int{{3, 3}, {5, -1}}, kClosest(points, 2))
	}
}

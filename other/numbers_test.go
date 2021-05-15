package other

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNumPairsDivisibleBy60(t *testing.T) {
	require.Equal(t, 3, numPairsDivisibleBy60([]int{60, 60, 60}))

	require.Equal(t, 1, numPairsDivisibleBy60([]int{10, 50, 60}))

	require.Equal(t, 3, numPairsDivisibleBy60([]int{30, 20, 150, 100, 40}))

	require.Equal(t, 1, numPairsDivisibleBy60([]int{20, 30, 40}))
}

func TestCalPoints(t *testing.T) {
	{
		ops := []string{"5", "2", "C", "D", "+"}
		require.Equal(t, 30, calPoints(ops))
	}
	{
		ops := []string{"5", "-2", "4", "C", "D", "9", "+", "+"}
		require.Equal(t, 27, calPoints(ops))
	}
}

func TestThreeSum(t *testing.T) {
	{
		nums := []int{1}
		require.Equal(t, [][]int{}, threeSum(nums))
	}
	{
		nums := []int{1, 2}
		require.Equal(t, [][]int{}, threeSum(nums))
	}
	{
		nums := []int{1, 2, 3}
		require.Equal(t, [][]int{}, threeSum(nums))
	}
	{
		nums := []int{1, 2, -3}
		require.Equal(t, [][]int{{-3, 1, 2}}, threeSum(nums))
	}
	{
		nums := []int{0, 0, 0, 0, 0}
		require.Equal(t, [][]int{{0, 0, 0}}, threeSum(nums))
	}
	{
		nums := []int{0, 0, 1, -1}
		require.Equal(t, [][]int{{-1, 0, 1}}, threeSum(nums))
	}
	{
		nums := []int{0, 3, -2, -3, -1}
		require.Equal(t, [][]int{{-3, 0, 3}, {-2, -1, 3}}, threeSum(nums))
	}
	{
		nums := []int{-1, 0, 1, 2, -1, -4}
		result := threeSum(nums)
		sort.Slice(result, func(i, j int) bool {
			key1 := result[i][0]*100 + result[i][1]*10 + result[i][2]
			key2 := result[j][0]*100 + result[j][1]*10 + result[j][2]
			return key1 < key2
		})
		require.Equal(t, [][]int{{-1, -1, 2}, {-1, 0, 1}}, result)
	}
	{
		nums := []int{-4, -2, 2, 0, 4, 1, 1, 3, -1, -1}
		result := threeSum(nums)
		sort.Slice(result, func(i, j int) bool {
			key1 := result[i][0]*100 + result[i][1]*10 + result[i][2]
			key2 := result[j][0]*100 + result[j][1]*10 + result[j][2]
			return key1 < key2
		})
		require.Equal(t, [][]int{{-4, 0, 4}, {-4, 1, 3}, {-2, -1, 3}, {-2, 0, 2}, {-2, 1, 1}, {-1, -1, 2}, {-1, 0, 1}}, result)
	}
	{
		nums := []int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}
		result := threeSum(nums)
		sort.Slice(result, func(i, j int) bool {
			key1 := result[i][0]*100 + result[i][1]*10 + result[i][2]
			key2 := result[j][0]*100 + result[j][1]*10 + result[j][2]
			return key1 < key2
		})
		require.Equal(t, [][]int{{-4, 0, 4}, {-4, 1, 3}, {-3, -1, 4}, {-3, 0, 3}, {-3, 1, 2}, {-2, -1, 3}, {-2, 0, 2}, {-1, -1, 2}, {-1, 0, 1}}, result)
	}
}

func TestSqrt(t *testing.T) {
	require.Equal(t, 2, mySqrt(8))
	require.Equal(t, 2, mySqrt(4))
	require.Equal(t, 3, mySqrt(9))
	require.Equal(t, 16, mySqrt(260))
}

func TestMostPopularFraction(t *testing.T) {
	{
		up := []int{1, 2, 3, 4}
		down := []int{2, 4, 6, 5}
		require.Equal(t, 3, mostPopularFraction(up, down))
	}
	{
		up := []int{0, 0}
		down := []int{2, 4}
		require.Equal(t, 1, mostPopularFraction(up, down))
	}
	{
		up := []int{1, 2, 3}
		down := []int{1, 2, 7}
		require.Equal(t, 2, mostPopularFraction(up, down))
	}
	{
		up := []int{1, 5, 10, 4, 7}
		down := []int{2, 7, 14, 3, 5}
		require.Equal(t, 2, mostPopularFraction(up, down))
	}
}

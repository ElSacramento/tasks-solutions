package nodes

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShortestPathBinaryMatrix(t *testing.T) {
	{
		grid := [][]int{
			{0, 1},
			{1, 0},
		}
		require.Equal(t, 2, shortestPathBinaryMatrix(grid))
	}
	{
		grid := [][]int{
			{0, 0},
			{0, 0},
		}
		require.Equal(t, 2, shortestPathBinaryMatrix(grid))
	}
	{
		grid := [][]int{
			{1, 1},
			{1, 1},
		}
		require.Equal(t, -1, shortestPathBinaryMatrix(grid))
	}
	{
		grid := [][]int{
			{0, 0},
			{1, 0},
		}
		require.Equal(t, 2, shortestPathBinaryMatrix(grid))
	}
	{
		grid := [][]int{
			{0, 0},
			{0, 1},
		}
		require.Equal(t, -1, shortestPathBinaryMatrix(grid))
	}
	{
		grid := [][]int{
			{0, 0, 0},
			{1, 1, 0},
			{1, 1, 0},
		}
		require.Equal(t, 4, shortestPathBinaryMatrix(grid))
	}
	{
		grid := [][]int{
			{1, 0, 0},
			{1, 1, 0},
			{1, 1, 0},
		}
		require.Equal(t, -1, shortestPathBinaryMatrix(grid))
	}
	{
		grid := [][]int{
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0},
		}
		require.Equal(t, 4, shortestPathBinaryMatrix(grid))
	}
	{
		grid := [][]int{
			{0, 1, 1, 0, 0, 0},
			{0, 1, 0, 1, 1, 0},
			{0, 1, 1, 0, 1, 0},
			{0, 0, 0, 1, 1, 0},
			{1, 1, 1, 1, 1, 0},
			{1, 1, 1, 1, 1, 0},
		}
		require.Equal(t, 14, shortestPathBinaryMatrix(grid))
	}
}

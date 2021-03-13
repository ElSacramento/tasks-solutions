package nodes

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindTheCity(t *testing.T) {
	{
		edges := [][]int{
			{0, 1, 3},
			{1, 2, 1},
			{1, 4, 1},
			{1, 3, 5},
			{2, 3, 2},
			{2, 4, 1},
			{2, 5, 1},
		}
		require.Equal(t, 0, findTheCity(6, edges, 4))
	}
	{
		edges := [][]int{
			{0, 1, 2},
			{1, 2, 1},
			{1, 4, 1},
			{1, 3, 5},
			{2, 3, 2},
			{2, 4, 1},
			{2, 5, 1},
		}
		require.Equal(t, 3, findTheCity(6, edges, 2))
	}
	{
		edges := [][]int{
			{0, 1, 3},
			{1, 2, 1},
			{1, 3, 4},
			{2, 3, 1},
		}
		require.Equal(t, 3, findTheCity(4, edges, 4))
	}
	{
		edges := [][]int{
			{0, 1, 3},
			{1, 2, 1},
			{1, 3, 4},
			{2, 3, 1},
		}
		require.Equal(t, 0, findTheCity(4, edges, 3))
	}
	{
		edges := [][]int{
			{0, 1, 2},
			{0, 4, 8},
			{1, 2, 3},
			{1, 4, 2},
			{2, 3, 1},
			{3, 4, 1},
		}
		require.Equal(t, 0, findTheCity(5, edges, 2))
	}
	{
		edges := [][]int{
			{0, 1, 10},
			{0, 2, 1},
			{2, 3, 1},
			{1, 3, 1},
			{1, 4, 1},
			{4, 5, 10},
		}
		require.Equal(t, 5, findTheCity(6, edges, 20))
	}
}

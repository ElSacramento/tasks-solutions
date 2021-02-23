package unionfind

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNumberOfIslands(t *testing.T) {
	{
		land := [][]byte{
			{1, 1, 1, 1, 0},
			{1, 1, 0, 1, 0},
			{1, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
		}
		require.Equal(t, 1, numIslands(land))
	}
	{
		land := [][]byte{
			{1, 1, 0, 0},
			{1, 0, 0, 1},
			{0, 0, 1, 1},
		}
		require.Equal(t, 2, numIslands(land))
	}
	{
		land := [][]byte{
			{1, 1, 0},
			{1, 0, 0},
			{0, 1, 0},
		}
		require.Equal(t, 2, numIslands(land))
	}
	{
		land := [][]byte{
			{1, 1},
			{1, 1},
		}
		require.Equal(t, 1, numIslands(land))
	}
	{
		land := [][]byte{
			{0, 0},
			{0, 0},
		}
		require.Equal(t, 0, numIslands(land))
	}
	{
		land := [][]byte{
			{1, 0},
			{0, 1},
		}
		require.Equal(t, 2, numIslands(land))
	}
	{
		land := [][]byte{
			{0, 0, 0, 0},
			{0, 1, 0, 0},
			{0, 1, 0, 0},
			{1, 1, 1, 0},
			{0, 1, 0, 0},
		}
		require.Equal(t, 1, numIslands(land))
	}
}

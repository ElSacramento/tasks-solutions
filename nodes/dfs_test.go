package nodes

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindCircleNum(t *testing.T) {
	{
		m := [][]int{
			{1, 1, 0},
			{1, 1, 0},
			{0, 0, 1},
		}
		require.Equal(t, 2, findCircleNum(m))
	}
	{
		m := [][]int{
			{1, 1, 0},
			{1, 1, 1},
			{0, 1, 1},
		}
		require.Equal(t, 1, findCircleNum(m))
	}
	{
		m := [][]int{
			{0, 0, 0},
			{0, 0, 0},
			{0, 0, 0},
		}
		require.Equal(t, 3, findCircleNum(m))
	}
}

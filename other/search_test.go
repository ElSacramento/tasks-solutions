package other

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindPeakElement(t *testing.T) {
	require.Equal(t, 1, findPeakElement([]int{1, 2, 1, 0, -1, -2, 1, 0}))
	require.Equal(t, 0, findPeakElement([]int{3, 2, 1}))
	require.Equal(t, 1, findPeakElement([]int{1, 2}))
	require.Equal(t, 0, findPeakElement([]int{1}))
	require.Equal(t, 0, findPeakElement([]int{-2147483647, -2147483648}))
}

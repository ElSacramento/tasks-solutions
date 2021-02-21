package other

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNumPairsDivisibleBy60(t *testing.T) {
	require.Equal(t, 3, numPairsDivisibleBy60([]int{60, 60, 60}))

	require.Equal(t, 1, numPairsDivisibleBy60([]int{10, 50, 60}))

	require.Equal(t, 3, numPairsDivisibleBy60([]int{30, 20, 150, 100, 40}))

	require.Equal(t, 1, numPairsDivisibleBy60([]int{20, 30, 40}))
}

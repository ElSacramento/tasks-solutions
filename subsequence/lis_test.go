package subsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLIS(t *testing.T) {
	{
		arr := []int{10, 9, 2, 5, 3, 7, 101, 18}
		assert.Equal(t, 4, LengthOfLIS(arr))
	}
	{
		arr := []int{10, 4, 5, 1, 7}
		assert.Equal(t, 3, LengthOfLIS(arr))
	}
	{
		arr := []int{10, 4, 5, 1, 2, 3, 7}
		assert.Equal(t, 4, LengthOfLIS(arr))
	}
	{
		arr := []int{10, 10, 10}
		assert.Equal(t, 1, LengthOfLIS(arr))
	}
	{
		arr := []int{3, 2, 1}
		assert.Equal(t, 1, LengthOfLIS(arr))
	}
}

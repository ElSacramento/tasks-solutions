package window

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTotalFruit(t *testing.T) {
	{
		arr := []int{1, 2, 3, 2, 2}
		assert.Equal(t, 4, TotalFruit(arr))
	}
	{
		arr := []int{0, 1, 2, 2}
		assert.Equal(t, 3, TotalFruit(arr))
	}
	{
		arr := []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}
		assert.Equal(t, 5, TotalFruit(arr))
	}
	{
		arr := []int{1, 2, 3}
		assert.Equal(t, 2, TotalFruit(arr))
	}
	{
		arr := []int{1, 1, 1, 1}
		assert.Equal(t, 4, TotalFruit(arr))
	}
	{
		arr := []int{1, 1, 6, 5, 6, 6, 1, 1, 1, 1}
		assert.Equal(t, 6, TotalFruit(arr))
	}
	{
		arr := []int{1, 1, 2, 1, 2}
		assert.Equal(t, 5, TotalFruit(arr))
	}
	{
		arr := []int{1, 2, 1, 2, 3}
		assert.Equal(t, 4, TotalFruit(arr))
	}
}

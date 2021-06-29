package other

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNumTimesAllBlue(t *testing.T) {
	{
		light := []int{2, 3, 1, 4, 5}
		require.Equal(t, 3, numTimesAllBlue(light))
	}
	{
		light := []int{1, 2, 3}
		require.Equal(t, 3, numTimesAllBlue(light))
	}
	{
		light := []int{3, 2, 1}
		require.Equal(t, 1, numTimesAllBlue(light))
	}
	{
		light := []int{2, 1, 3, 5, 4}
		require.Equal(t, 3, numTimesAllBlue(light))
	}
	{
		light := []int{3, 2, 4, 1, 5}
		require.Equal(t, 2, numTimesAllBlue(light))
	}
	{
		light := []int{4, 1, 2, 3}
		require.Equal(t, 1, numTimesAllBlue(light))
	}
	{
		light := []int{2, 1, 4, 3, 6, 5}
		require.Equal(t, 3, numTimesAllBlue(light))
	}
}

func TestIsValidBrackets(t *testing.T) {
	{
		s := "([{}])"
		require.True(t, isValid(s))
	}
	{
		s := "()"
		require.True(t, isValid(s))
	}
	{
		s := "(){}[]"
		require.True(t, isValid(s))
	}
	{
		s := "(]"
		require.False(t, isValid(s))
	}
	{
		s := "([)]"
		require.False(t, isValid(s))
	}
	{
		s := "{{[[(())]]}}"
		require.True(t, isValid(s))
	}
	{
		s := "{[(])}"
		require.False(t, isValid(s))
	}
	{
		s := "("
		require.False(t, isValid(s))
	}
}
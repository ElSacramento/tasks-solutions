package subsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLongestSubstring(t *testing.T) {
	assert.Equal(t, 3, LongestSubstring("aaabb", 3))
	assert.Equal(t, 5, LongestSubstring("ababbc", 2))
	assert.Equal(t, 1, LongestSubstring("a", 1))
	assert.Equal(t, 0, LongestSubstring("a", 10))
	assert.Equal(t, 10, LongestSubstring("abaabcdcdaef", 2))
	assert.Equal(t, 4, LongestSubstring("babacbaba", 2))
	assert.Equal(t, 4, LongestSubstring("aadbabacdaab", 2))
}

func TestLongestCommonPrefix(t *testing.T) {
	{
		strs := []string{"flower", "flow", "flight"}
		require.Equal(t, "fl", longestCommonPrefix(strs))
	}
	{
		strs := []string{"abc", "ab", "ad"}
		require.Equal(t, "a", longestCommonPrefix(strs))
	}
	{
		strs := []string{"flower", "flow"}
		require.Equal(t, "flow", longestCommonPrefix(strs))
	}
	{
		strs := []string{"abcd", "abcde", "abcf"}
		require.Equal(t, "abc", longestCommonPrefix(strs))
	}
}

package other

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBreakPalindrome(t *testing.T) {
	require.Equal(t, "aab", breakPalindrome("aaa"))
	require.Equal(t, "acb", breakPalindrome("aca"))
	require.Equal(t, "aacba", breakPalindrome("abcba"))
	require.Equal(t, "aaddca", breakPalindrome("acddca"))
	require.Equal(t, "acce", breakPalindrome("ecce"))
}

func TestLengthOfLongestSubstring(t *testing.T) {
	require.Equal(t, 3, lengthOfLongestSubstring("abcabcbb"))
	require.Equal(t, 6, lengthOfLongestSubstring("#x Ab9"))
	require.Equal(t, 1, lengthOfLongestSubstring("bbbbb"))
	require.Equal(t, 3, lengthOfLongestSubstring("pwwkew"))
	require.Equal(t, 4, lengthOfLongestSubstring("abcda"))
	require.Equal(t, 3, lengthOfLongestSubstring("abcba"))
	require.Equal(t, 0, lengthOfLongestSubstring(""))
}
